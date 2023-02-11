package socket

import (
	"github.com/gorilla/websocket"
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

type Session struct {
	conn        *websocket.Conn
	rawSendChan chan string
}

type ConnectManager struct {
	sessionMap     map[net.Addr]*Session
	sessionMapLock sync.RWMutex
}

func (c *ConnectManager) Close() {
	c.closeAllSession()
}

// acceptHandler 接受并创建会话的处理函数
func (c *ConnectManager) acceptHandler(w http.ResponseWriter, r *http.Request) {
	// 连接提升至websocket
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	logger.Info("client connect, address: %v", conn.RemoteAddr())
	session := &Session{
		conn:        conn,
		rawSendChan: make(chan string, 1000),
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
	conn := session.conn
	for {
		_ = conn.SetReadDeadline(time.Now().Add(time.Second * ConnRecvTimeout))
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			logger.Error("exit session recv loop, address: %v", conn.RemoteAddr())
			c.closeSession(session)
			break
		}
		logger.Info("messageType: %v, data: %v", messageType, string(data))
	}
}

// sendHandler 发送消息的处理函数
func (c *ConnectManager) sendHandler(session *Session) {
	conn := session.conn
	for {
		_ = conn.SetWriteDeadline(time.Now().Add(time.Second * ConnSendTimeout))
		// err := conn.WriteMessage(1, []byte("测试测试~"))
		// if err != nil {
		// 	logger.Error("exit session send loop, addr: %v", conn.RemoteAddr())
		// 	c.closeSession(session)
		// 	break
		// }
	}
}

// closeSession 关闭某一会话
func (c *ConnectManager) closeSession(session *Session) {
	conn := session.conn
	c.sessionMapLock.RLock()
	delete(c.sessionMap, conn.RemoteAddr())
	c.sessionMapLock.RUnlock()
	_ = conn.Close()
}

// closeAllSession 关闭全部会话
func (c *ConnectManager) closeAllSession() {
	c.sessionMapLock.RLock()
	for _, s := range c.sessionMap {
		c.closeSession(s)
	}
	c.sessionMapLock.RUnlock()
}

func NewConnectManager() *ConnectManager {
	connectManager := new(ConnectManager)
	connectManager.sessionMap = make(map[net.Addr]*Session, 100)

	http.HandleFunc("/", connectManager.acceptHandler)
	http.HandleFunc("/api", connectManager.acceptHandler)
	go func() {
		err := http.ListenAndServe("0.0.0.0:8080", nil)
		if err != nil {
			logger.Error("http listen error: %v", err)
			return
		}
	}()

	return connectManager
}
