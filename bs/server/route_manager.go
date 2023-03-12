package server

import (
	"moon-bot/bs/bot"
	"moon-bot/common/constant"
	"moon-bot/common/mq"
	"moon-bot/pkg/logger"
	"moon-bot/protocol/cmd"
	"moon-bot/protocol/pb"

	"github.com/golang/protobuf/proto"
)

// RouteManager 路由管理器
type RouteManager struct {
	routeFuncMap map[string]bot.RouteFunc // 路由函数注册
}

func NewRouteManager() *RouteManager {
	r := new(RouteManager)
	r.routeFuncMap = map[string]bot.RouteFunc{
		cmd.MessageEvent: r.MessageEvent,
	}

	return r
}

// doRoute 执行路由
func (r *RouteManager) doRoute(cmdName string, userId int64, payloadMsg proto.Message) {
	routeFunc, ok := r.routeFuncMap[cmdName]
	if !ok {
		logger.Error("no route for msg, cmdName: %v", cmdName)
		return
	}
	// 转发路由处理函数到bot
	bot.GetManageBot().GetRouteMsgChan() <- &bot.RouteMsg{
		UserId:     userId,
		RouteFunc:  routeFunc,
		PayloadMsg: payloadMsg,
	}
}

// handleNetMsg 处理网络消息
func (r *RouteManager) handleNetMsg(netMsg *mq.NetMsg) {
	switch netMsg.MsgType {
	case mq.MsgTypeProto:
		protoMsg := netMsg.ProtoMsg
		// 确保消息为生命周期连接
		if protoMsg.CmdName == cmd.MetaEvent && protoMsg.PayloadMessage.(*pb.MetaEvent).SubType == constant.MetaEventTypeLifecycleSubTypeConnect {
			// 用户连接
			bot.GetManageBot().GetUserCtrlMsgChan() <- &bot.UserCtrlMsg{
				UserCtrlType: bot.UserCtrlTypeConnect,
				UserId:       protoMsg.UserId,
			}
			return
		}
		r.doRoute(protoMsg.CmdName, protoMsg.UserId, protoMsg.PayloadMessage)
	case mq.MsgTypeOffline:
		// 用户离线消息
		offlineMsg := netMsg.OfflineMsg
		bot.GetManageBot().GetUserCtrlMsgChan() <- &bot.UserCtrlMsg{
			UserCtrlType: bot.UserCtrlTypeDisconnect,
			UserId:       offlineMsg.UserId,
		}
	}
}
