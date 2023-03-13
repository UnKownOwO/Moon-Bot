package server

import (
	"moon-bot/bs/bot"
	"moon-bot/pkg/logger"
	"moon-bot/protocol/cmd"

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
