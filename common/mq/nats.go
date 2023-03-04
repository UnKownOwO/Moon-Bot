package mq

import (
	"moon-bot/pkg/logger"
	"moon-bot/protocol/cmd"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/vmihailenco/msgpack"
)

type ServerType string

const (
	ServerTypeGate ServerType = "gate" // 网关
	ServerTypeBs   ServerType = "bs"   // bot server
)

type MessageQueue struct {
	natsConn         *nats.Conn
	serverType       string
	natsMsgChan      chan *nats.Msg
	netMsgInputChan  chan *NetMsg
	netMsgOutputChan chan *NetMsg
}

func NewMessageQueue(serverType ServerType) *MessageQueue {
	messageQueue := new(MessageQueue)
	messageQueue.serverType = string(serverType)

	// 连接nats
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		logger.Error("nats connect error: %v", err)
		return nil
	}
	messageQueue.natsConn = nc

	// 通道订阅
	messageQueue.natsMsgChan = make(chan *nats.Msg, 1000)
	_, err = nc.ChanSubscribe(messageQueue.serverType, messageQueue.natsMsgChan)
	if err != nil {
		logger.Error("nats subscribe error: %v", err)
		return nil
	}

	messageQueue.netMsgInputChan = make(chan *NetMsg, 1000)
	messageQueue.netMsgOutputChan = make(chan *NetMsg, 1000)

	// 处理协程
	go messageQueue.recvHandler()
	go messageQueue.sendHandler()

	return messageQueue
}

// recvHandler 处理消息接收
func (m *MessageQueue) recvHandler() {
	for {
		natsMsg := <-m.natsMsgChan
		// 使用msgpack反序列化数据
		netMsg := new(NetMsg)
		err := msgpack.Unmarshal(natsMsg.Data, netMsg)
		if err != nil {
			logger.Error("unmarshal net msg error: %v", err)
			return
		}
		// 如果来自gate则反序列化payload
		if netMsg.MsgType == MsgTypeGate {
			payloadMessage := cmd.GetCmdObjManager().GetProtoObjByCmdName(netMsg.GateMsg.CmdName)
			err := proto.Unmarshal(netMsg.GateMsg.PayloadMessageData, payloadMessage)
			if err != nil {
				logger.Error("unmarshal gate msg payload message error: %v", err)
				return
			}
			netMsg.GateMsg.PayloadMessage = payloadMessage
		}
		m.netMsgOutputChan <- netMsg
	}
}

// sendHandler 处理消息发送
func (m *MessageQueue) sendHandler() {
	for {
		netMsg := <-m.netMsgInputChan
		// 如果来自gate则序列化payload
		if netMsg.MsgType == MsgTypeGate {
			payloadMessageData, err := proto.Marshal(netMsg.GateMsg.PayloadMessage)
			if err != nil {
				logger.Error("marshal gate msg payload message error: %v", err)
				return
			}
			netMsg.GateMsg.PayloadMessageData = payloadMessageData
		}
		// 使用msgpack序列化数据
		netMsgData, err := msgpack.Marshal(netMsg)
		if err != nil {
			logger.Error("marshal net msg error: %v", netMsgData)
			return
		}
		// 通过nats发送数据
		natsMsg := nats.NewMsg(netMsg.TargetServer)
		natsMsg.Data = netMsgData
		err = m.natsConn.PublishMsg(natsMsg)
		if err != nil {
			logger.Error("publish nats msg error: %v", err)
			return
		}
	}
}

// GetNetMsgChan 获取输出消息的管道
func (m *MessageQueue) GetNetMsgChan() <-chan *NetMsg {
	return m.netMsgOutputChan
}

// Send 发送消息
func (m *MessageQueue) Send(targetServer ServerType, netMsg *NetMsg) {
	netMsg.TargetServer = string(targetServer)
	m.netMsgInputChan <- netMsg
}

// Close 关闭消息队列
func (m *MessageQueue) Close() {
	m.natsConn.Close()
}
