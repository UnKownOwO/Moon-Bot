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
	conn        *websocket.Conn
	state       SessionState       // 会话状态
	userId      int64              // 机器人QQ号
	sendRawChan chan *ProtoMessage // 发送proto消息
}

// handleMessage 处理会话消息
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
		// 处理会话生命周期
		c.sessionLifecycle(session, metaEvent)
	}
	if session.state != SessionStateActive && message.CmdName != cmd.MetaEvent {
		logger.Error("session not active packet drop, cmdName: %v, address: %v", message.CmdName, session.conn.RemoteAddr())
		return
	}
	// 转发消息到bs
	protoMsg := new(mq.ProtoMsg)
	protoMsg.UserId = session.userId
	protoMsg.CmdName = message.CmdName
	protoMsg.PayloadMessage = message.PayloadMessage
	mq.Send(mq.ServerTypeBs, &mq.NetMsg{
		MsgType:  mq.MsgTypeProto,
		ProtoMsg: protoMsg,
	})
}

// sessionLifecycle 处理会话生命周期
func (c *ConnectManager) sessionLifecycle(session *Session, msg *pb.MetaEvent) {
	// 关联session信息 不然包发不出去
	session.userId = msg.SelfId
	c.AddSession(session)
	c.createSessionChan <- session
	// 设置会话为已连接
	session.state = SessionStateActive
}
