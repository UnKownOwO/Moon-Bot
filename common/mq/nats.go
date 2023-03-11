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
	ServerTypeBs   ServerType = "bs"   // server server
)

type MessageQueue struct {
	natsConn         *nats.Conn
	serverType       string
	natsMsgChan      chan *nats.Msg
	netMsgInputChan  chan *NetMsg
	netMsgOutputChan chan *NetMsg
}

var messageQueue *MessageQueue

func InitMessageQueue(serverType ServerType) {
	if messageQueue == nil {
		messageQueue = new(MessageQueue)
		messageQueue.serverType = string(serverType)

		// 连接nats
		nc, err := nats.Connect(nats.DefaultURL)
		if err != nil {
			logger.Error("nats connect error: %v", err)
			return
		}
		messageQueue.natsConn = nc

		// 通道订阅
		messageQueue.natsMsgChan = make(chan *nats.Msg, 1000)
		_, err = nc.ChanSubscribe(messageQueue.serverType, messageQueue.natsMsgChan)
		if err != nil {
			logger.Error("nats subscribe error: %v", err)
			return
		}

		messageQueue.netMsgInputChan = make(chan *NetMsg, 1000)
		messageQueue.netMsgOutputChan = make(chan *NetMsg, 1000)

		// 处理协程
		go messageQueue.recvHandler()
		go messageQueue.sendHandler()
	}
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
			continue
		}
		// 如果为proto类型则反序列化payload
		if netMsg.MsgType == MsgTypeProto {
			payloadMessage := cmd.GetCmdObjManager().GetProtoObjByCmdName(netMsg.ProtoMsg.CmdName)
			if payloadMessage == nil {
				continue
			}
			err := proto.Unmarshal(netMsg.ProtoMsg.PayloadMessageData, payloadMessage)
			if err != nil {
				logger.Error("unmarshal payload msg error: %v", err)
				continue
			}
			netMsg.ProtoMsg.PayloadMessage = payloadMessage
		}
		m.netMsgOutputChan <- netMsg
	}
}

// sendHandler 处理消息发送
func (m *MessageQueue) sendHandler() {
	for {
		netMsg := <-m.netMsgInputChan
		// 如果为proto类型则序列化payload
		if netMsg.MsgType == MsgTypeProto {
			payloadMessageData, err := proto.Marshal(netMsg.ProtoMsg.PayloadMessage)
			if err != nil {
				logger.Error("marshal payload msg error: %v", err)
				continue
			}
			netMsg.ProtoMsg.PayloadMessageData = payloadMessageData
		}
		// 使用msgpack序列化数据
		netMsgData, err := msgpack.Marshal(netMsg)
		if err != nil {
			logger.Error("marshal net msg error: %v", netMsgData)
			continue
		}
		// 通过nats发送数据
		natsMsg := nats.NewMsg(netMsg.TargetServer)
		natsMsg.Data = netMsgData
		err = m.natsConn.PublishMsg(natsMsg)
		if err != nil {
			logger.Error("publish nats msg error: %v", err)
			continue
		}
	}
}

// GetNetMsgChan 获取输出消息的管道
func GetNetMsgChan() <-chan *NetMsg {
	return messageQueue.netMsgOutputChan
}

// Send 发送消息
func Send(targetServer ServerType, netMsg *NetMsg) {
	netMsg.TargetServer = string(targetServer)
	messageQueue.netMsgInputChan <- netMsg
}

// Close 关闭消息队列
func Close() {
	messageQueue.natsConn.Close()
}
