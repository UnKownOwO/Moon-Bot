package net

import (
	"moon-bot/common/constant"
	"moon-bot/common/mq"
	"moon-bot/pkg/logger"
	"moon-bot/protocol/cmd"
	"moon-bot/protocol/pb"

	"github.com/gorilla/websocket"
)

type SessionState uint8

const (
	SessionStateWait   SessionState = iota // 等待连接
	SessionStateActive                     // 活跃状态
	SessionStateClose                      // 关闭状态
)

type Session struct {
	conn    *websocket.Conn
	state   SessionState
	account int64
}

// handleMessage 处理消息
func (c *ConnectManager) handleMessage(session *Session, message *ProtoMessage) {
	switch message.CmdName {
	case cmd.MetaEvent:
		// 确保用户未连接
		if session.state != SessionStateWait {
			return
		}
		metaEvent := message.PayloadMessage.(*pb.MetaEvent)
		// 确保为生命周期事件
		if metaEvent.MetaEventType != constant.MetaEventTypeLifecycle {
			return
		}
		// 设置会话为已连接
		session.account = metaEvent.SelfId
		session.state = SessionStateActive
	}
	if session.state != SessionStateActive && message.CmdName != cmd.MetaEvent {
		logger.Error("session not active packet drop, cmdName: %v, address: %v", message.CmdName, session.conn.RemoteAddr())
		return
	}
	// 转发消息到BS
	gateMsg := new(mq.GateMsg)
	gateMsg.Account = session.account
	gateMsg.CmdName = message.CmdName
	gateMsg.PayloadMessage = message.PayloadMessage
	c.messageQueue.Send(mq.ServerTypeBs, &mq.NetMsg{
		MsgType: mq.MsgTypeGate,
		GateMsg: gateMsg,
	})
}
