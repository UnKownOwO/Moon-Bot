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
	r.routeFuncMap = make(map[string]bot.RouteFunc)
	// 初始化所有路由
	r.initRoute()

	return r
}

// initRoute 初始化路由
func (m *RouteManager) initRoute() {
	m.regRoute(cmd.MessageEvent, m.MessageEvent)
}

// regRoute 注册路由
func (m *RouteManager) regRoute(cmdName string, routeFunc bot.RouteFunc) {
	m.routeFuncMap[cmdName] = routeFunc
}

// doRoute 执行路由
func (m *RouteManager) doRoute(cmdName string, userId int64, payloadMsg proto.Message) {
	routeFunc, ok := m.routeFuncMap[cmdName]
	if !ok {
		logger.Error("no route for msg, cmdName: %v", cmdName)
		return
	}
	bot.GetManageBot().GetRouteMsgChan() <- &bot.RouteMsg{
		UserId:     userId,
		RouteFunc:  routeFunc,
		PayloadMsg: payloadMsg,
	}
}

// HandleNetMsg 处理消息
func (m *RouteManager) HandleNetMsg(netMsg *mq.NetMsg) {
	switch netMsg.MsgType {
	case mq.MsgTypeProto:
		protoMsg := netMsg.ProtoMsg
		if protoMsg.CmdName == cmd.MetaEvent {
			metaEvent := protoMsg.PayloadMessage.(*pb.MetaEvent)
			if metaEvent.SubType == constant.MetaEventTypeLifecycleSubTypeConnect {
				// 用户连接
				bot.GetManageBot().GetConnectChan() <- metaEvent.SelfId
				return
			}
		}
		m.doRoute(protoMsg.CmdName, protoMsg.UserId, protoMsg.PayloadMessage)
	}
}
