package mq

import (
	"github.com/nats-io/nats.go"
	"moon-bot/pkg/logger"
)

const (
	ServerType_Gate string = "gate" // 网关
	ServerType_Bs   string = "bs"   // bot server
)

type MessageQueue struct {
	natsConn    *nats.Conn
	natsMsgChan chan *nats.Msg
}

// Close 关闭消息队列
func (m *MessageQueue) Close() {
	m.natsConn.Close()
}

// recvHandler 处理消息接收
func (m *MessageQueue) recvHandler() {

}

// sendHandler 处理消息发送
func (m *MessageQueue) sendHandler() {

}

func NewMessageQueue(serverType string) {
	messageQueue := new(MessageQueue)

	// 连接nats
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		logger.Error("nats connect error: %v", err)
		return
	}
	messageQueue.natsConn = nc

	// 通道订阅
	messageQueue.natsMsgChan = make(chan *nats.Msg, 1000)
	_, err = nc.ChanSubscribe(serverType, messageQueue.natsMsgChan)
	if err != nil {
		logger.Error("nats subscribe error: %v", err)
		return
	}

	// 处理协程
	go messageQueue.recvHandler()
	go messageQueue.sendHandler()
}
