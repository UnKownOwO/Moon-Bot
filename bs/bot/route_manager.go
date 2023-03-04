package bot

import (
	"moon-bot/bs/model"
	"moon-bot/common/mq"
	"moon-bot/pkg/logger"
	"moon-bot/protocol/cmd"

	"github.com/golang/protobuf/proto"
)

type HandlerFunc func(bot *model.Bot, payloadMsg proto.Message)

type RouteManager struct {
	handlerFuncRouteMap map[string]HandlerFunc
}

func NewRouteManager() *RouteManager {
	r := new(RouteManager)
	r.handlerFuncRouteMap = make(map[string]HandlerFunc)

	// 初始化所有路由
	r.initRoute()

	return r
}

// initRoute 初始化路由
func (r *RouteManager) initRoute() {
	r.regRoute(cmd.MessageEvent, botManager.messageEvent)
}

// regRoute 注册路由
func (r *RouteManager) regRoute(cmdName string, handlerFunc HandlerFunc) {
	r.handlerFuncRouteMap[cmdName] = handlerFunc
}

// doRoute 执行路由
func (r *RouteManager) doRoute(cmdName string, account int64, payloadMsg proto.Message) {
	logger.Debug("message: %v", payloadMsg)
	handlerFunc, ok := r.handlerFuncRouteMap[cmdName]
	if !ok {
		logger.Error("no route for msg, cmdName: %v", cmdName)
		return
	}
	handlerFunc(&model.Bot{}, payloadMsg)
}

// HandleNetMsg 处理消息
func (r *RouteManager) HandleNetMsg(netMsg *mq.NetMsg) {
	switch netMsg.MsgType {
	case mq.MsgTypeGate:
		gateMsg := netMsg.GateMsg
		r.doRoute(gateMsg.CmdName, gateMsg.Account, gateMsg.PayloadMessage)
	}
}
