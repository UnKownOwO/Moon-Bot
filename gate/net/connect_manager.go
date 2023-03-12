package net

import (
	"moon-bot/common/mq"
	"moon-bot/pkg/logger"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	ConnRecvTimeout = 30 // 收包超时时间 秒
)

// 使用websocket默认配置
var upGrader = websocket.Upgrader{
	// 校验来源
	CheckOrigin: func(r *http.Request) bool {
		// if r.Header.Get("Origin") != "http://"+"0.0.0.0" {
		// 	return false
		// }
		return true
	},
}

type ConnectManager struct {
	// 会话
	sessionAddrMap     map[net.Addr]*Session
	sessionUserIdMap   map[int64]*Session
	sessionMapLock     sync.RWMutex
	createSessionChan  chan *Session
	destroySessionChan chan *Session
}

func NewConnectManager() *ConnectManager {
	r := new(ConnectManager)
	r.sessionAddrMap = make(map[net.Addr]*Session)
	r.sessionUserIdMap = make(map[int64]*Session)
	r.createSessionChan = make(chan *Session, 100)
	r.destroySessionChan = make(chan *Session, 100)

	go r.listenServe("0.0.0.0:8080")
	go r.recvMsgHandler()
	return r
}

func (c *ConnectManager) Close() {
	c.closeAllConn()
}

// recvMsgHandler 接收mq消息并处理
func (c *ConnectManager) recvMsgHandler() {
	logger.Debug("recv msg handler start")
	// 函数栈内缓存 添加删除事件走chan 避免频繁加锁
	userIdSessionMap := make(map[int64]*Session)
	for {
		select {
		case session := <-c.createSessionChan:
			// 创建函数栈内缓存的会话
			userIdSessionMap[session.userId] = session
		case session := <-c.destroySessionChan:
			// 删除函数栈内缓存的会话
			delete(userIdSessionMap, session.userId)
			close(session.sendRawChan)
		case netMsg := <-mq.GetNetMsgChan():
			// 接收bs的消息
			switch netMsg.MsgType {
			case mq.MsgTypeProto:
				// 发送消息给客户端
				protoMsg := netMsg.ProtoMsg
				session, ok := userIdSessionMap[protoMsg.UserId]
				if !ok {
					logger.Error("session not exist, userId: %v", protoMsg.UserId)
					return
				}
				session.sendRawChan <- &ProtoMessage{
					CmdName:        protoMsg.CmdName,
					PayloadMessage: protoMsg.PayloadMessage,
				}
			case mq.MsgTypeOffline:
				// 踢出用户
				offlineMsg := netMsg.OfflineMsg
				session, ok := userIdSessionMap[offlineMsg.UserId]
				if !ok {
					logger.Error("session not exist, userId: %v", offlineMsg.UserId)
					return
				}
				// 关闭用户的连接
				c.closeConn(session)
			}
		}
	}
}

// listenServe 监听服务
func (c *ConnectManager) listenServe(addr string) {
	logger.Debug("listen serve start")

	http.HandleFunc("/", c.handleAccept)
	http.HandleFunc("/api", c.handleAccept)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		logger.Error("http listen error: %v", err)
		return
	}
}

// handleAccept 接受并创建会话的处理函数
func (c *ConnectManager) handleAccept(w http.ResponseWriter, r *http.Request) {
	// 连接提升至websocket
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	logger.Info("client connect, addr: %v", conn.RemoteAddr())
	session := &Session{
		conn:        conn,
		state:       SessionStateWait,
		userId:      0,
		sendRawChan: make(chan *ProtoMessage, 100),
	}
	// 开始收发会话消息
	go c.recvHandler(session)
	go c.sendHandler(session)
}

// recvHandler 接收会话消息
func (c *ConnectManager) recvHandler(session *Session) {
	for {
		_ = session.conn.SetReadDeadline(time.Now().Add(time.Second * ConnRecvTimeout))
		_, data, err := session.conn.ReadMessage()
		if err != nil {
			logger.Error("exit recv loop, conn read error: %v, addr: %v", err, session.conn.RemoteAddr())
			c.closeConn(session)
			break
		}
		// 转换json为proto消息
		protoMessage := c.eventJsonToProtoMessage(data)
		if protoMessage == nil {
			continue
		}
		logger.Debug("[RECV] cmdName: %v, userId: %v, addr: %v, data: %v", protoMessage.CmdName, session.userId, session.conn.RemoteAddr(), protoMessage.PayloadMessage)
		// 处理消息
		c.handleMessage(session, protoMessage)
	}
}

// sendHandler 发送会话消息
func (c *ConnectManager) sendHandler(session *Session) {
	for {
		// 接收要发送的protoMsg
		protoMessage, ok := <-session.sendRawChan
		if !ok {
			logger.Error("exit send loop, send chan close, addr: %v", session.conn.RemoteAddr())
			c.closeConn(session)
			break
		}
		// 转换为json数据
		jsonData := c.protoMessageToApiJson(protoMessage)
		logger.Debug("[SEND] cmdName: %v, userId: %v, addr: %v, data: %v", protoMessage.CmdName, session.userId, session.conn.RemoteAddr(), string(jsonData))
		// 发送数据
		err := session.conn.WriteMessage(1, jsonData)
		if err != nil {
			logger.Error("exit send loop, conn write error: %v, addr: %v", err, session.conn.RemoteAddr())
			c.closeConn(session)
			break
		}
	}
}

// closeConn 关闭连接
func (c *ConnectManager) closeConn(session *Session) {
	if session == nil {
		return
	}
	// 会话已被关闭就没必要再关了
	if session.state == SessionStateClose {
		return
	}
	session.state = SessionStateClose
	// 通知bs用户下线
	if session.userId != 0 {
		mq.Send(mq.ServerTypeBs, &mq.NetMsg{
			MsgType: mq.MsgTypeOffline,
			OfflineMsg: &mq.OfflineMsg{
				UserId: session.userId,
			},
		})
	}
	// 删除会话
	c.DelSession(session)
	// 关闭连接
	_ = session.conn.Close()
	c.destroySessionChan <- session
}

// closeAllConn 关闭所有连接
func (c *ConnectManager) closeAllConn() {
	sessionList := make([]*Session, len(c.sessionAddrMap))
	c.sessionMapLock.RLock()
	for _, session := range c.sessionAddrMap {
		sessionList = append(sessionList, session)
	}
	c.sessionMapLock.RUnlock()
	for _, session := range sessionList {
		c.closeConn(session)
	}
}

// GetSessionByAddr 通过addr获取会话
func (c *ConnectManager) GetSessionByAddr(addr net.Addr) *Session {
	c.sessionMapLock.RLock()
	session := c.sessionAddrMap[addr]
	c.sessionMapLock.RUnlock()
	return session
}

// GetSessionByUserId 通过userId获取会话
func (c *ConnectManager) GetSessionByUserId(userId int64) *Session {
	c.sessionMapLock.RLock()
	session := c.sessionUserIdMap[userId]
	c.sessionMapLock.RUnlock()
	return session
}

// AddSession 添加会话
func (c *ConnectManager) AddSession(session *Session) {
	c.sessionMapLock.Lock()
	c.sessionAddrMap[session.conn.RemoteAddr()] = session
	c.sessionUserIdMap[session.userId] = session
	c.sessionMapLock.Unlock()
}

// DelSession 删除会话
func (c *ConnectManager) DelSession(session *Session) {
	c.sessionMapLock.Lock()
	delete(c.sessionAddrMap, session.conn.RemoteAddr())
	delete(c.sessionUserIdMap, session.userId)
	c.sessionMapLock.Unlock()
}
