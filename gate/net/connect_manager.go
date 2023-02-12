package net

import (
	"github.com/gorilla/websocket"
	"moon-bot/common/mq"
	"moon-bot/pkg/logger"
	"net"
	"net/http"
	"sync"
	"time"
)

const (
	ConnRecvTimeout = 30 // 收包超时时间 秒
	ConnSendTimeout = 10 // 发包超时时间 秒
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
	messageQueue   *mq.MessageQueue
	sessionMap     map[net.Addr]*Session
	sessionMapLock sync.RWMutex
}

func (c *ConnectManager) Close() {
	c.closeAllSession()
}

// handleAccept 接受并创建会话的处理函数
func (c *ConnectManager) handleAccept(w http.ResponseWriter, r *http.Request) {
	// 连接提升至websocket
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	logger.Info("client connect, address: %v", conn.RemoteAddr())
	session := &Session{
		conn: conn,
	}
	c.sessionMapLock.RLock()
	c.sessionMap[conn.RemoteAddr()] = session
	c.sessionMapLock.RUnlock()
	// 开始收发消息
	go c.recvHandler(session)
	go c.sendHandler(session)
}

// recvHandler 接收消息的处理函数
func (c *ConnectManager) recvHandler(session *Session) {
	for {
		_ = session.conn.SetReadDeadline(time.Now().Add(time.Second * ConnRecvTimeout))
		_, data, err := session.conn.ReadMessage()
		if err != nil {
			logger.Error("exit session recv loop, address: %v", session.conn.RemoteAddr())
			c.closeSession(session)
			break
		}
		// 转换json为proto消息
		protoMessage := c.eventJsonToProtoMessage(data)
		if protoMessage == nil {
			continue
		}
		logger.Debug("[RECV] cmdName: %v, account: %v, address: %v, data: %v", protoMessage.CmdName, session.account, session.conn.RemoteAddr(), protoMessage.PayloadMessage)
		// 处理消息
		c.handleMessage(session, protoMessage)
	}
}

// sendHandler 发送消息的处理函数
func (c *ConnectManager) sendHandler(session *Session) {
	for {
		_ = session.conn.SetWriteDeadline(time.Now().Add(time.Second * ConnSendTimeout))
		// err := conn.WriteMessage(1, []byte("测试测试~"))
		// if err != nil {
		// 	logger.Error("exit session send loop, addr: %v", conn.RemoteAddr())
		// 	c.closeSession(session)
		// 	break
		// }
		time.Sleep(time.Second)
	}
}

// closeSession 关闭某一会话
func (c *ConnectManager) closeSession(session *Session) {
	session.state = SessionStateClose
	c.sessionMapLock.RLock()
	delete(c.sessionMap, session.conn.RemoteAddr())
	c.sessionMapLock.RUnlock()
	_ = session.conn.Close()
}

// closeAllSession 关闭全部会话
func (c *ConnectManager) closeAllSession() {
	c.sessionMapLock.RLock()
	for _, s := range c.sessionMap {
		c.closeSession(s)
	}
	c.sessionMapLock.RUnlock()
}

func NewConnectManager(messageQueue *mq.MessageQueue) *ConnectManager {
	connectManager := new(ConnectManager)
	connectManager.messageQueue = messageQueue
	connectManager.sessionMap = make(map[net.Addr]*Session, 100)

	http.HandleFunc("/", connectManager.handleAccept)
	http.HandleFunc("/api", connectManager.handleAccept)
	go func() {
		err := http.ListenAndServe("0.0.0.0:8080", nil)
		if err != nil {
			logger.Error("http listen error: %v", err)
			return
		}
	}()

	return connectManager
}
