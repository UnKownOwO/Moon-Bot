package server

import (
	"moon-bot/bs/bot"
	"moon-bot/bs/event"
	"moon-bot/common/constant"
	"moon-bot/pkg/logger"
	"moon-bot/protocol/pb"

	"github.com/golang/protobuf/proto"
)

// MessageEvent 接收到消息
func (m *RouteManager) MessageEvent(bot *bot.Bot, payloadMsg proto.Message) {
	logger.Debug("消息事件收到准备通知模块")
	msg := payloadMsg.(*pb.MessageEvent)

	messageEvent := &event.MessageEvent{
		SubType:    msg.SubType,
		MessageId:  msg.MessageId,
		UserId:     msg.UserId,
		Message:    msg.Message,
		RawMessage: msg.RawMessage,
		Font:       msg.Font,
		Sender:     msg.Sender,
	}
	switch msg.MessageType {
	case constant.MessageEventTypePrivate:
		// 私聊消息
		moduleEvent := &event.PrivateMessageEvent{
			MessageEvent: messageEvent,
			TargetId:     msg.TargetId,
			TempSource:   msg.TempSource,
		}
		// 通知模块处理事件
		eventManager.HandleModuleEvent(event.ModuleEventIdPrivateMessage, bot, moduleEvent)
	case constant.MetaEventTypeGroup:
		// 群消息
		moduleEvent := &event.GroupMessageEvent{
			MessageEvent: messageEvent,
			GroupId:      msg.GroupId,
			Anonymous:    msg.Anonymous,
		}
		// 通知模块处理事件
		eventManager.HandleModuleEvent(event.ModuleEventIdPrivateMessage, bot, moduleEvent)
	}
}
