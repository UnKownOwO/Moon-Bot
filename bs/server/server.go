package server

import (
	"moon-bot/bs/bot"
	"moon-bot/bs/module"
	"moon-bot/common/mq"
	"moon-bot/pkg/logger"
)

var (
	routeManager *RouteManager // 路由管理器
	eventManager *EventManager // 事件管理器
)

type BotServer struct {
}

func NewBotServer() *BotServer {
	botServer := new(BotServer)
	// 初始化服务器管理器
	routeManager = NewRouteManager()
	eventManager = NewEventManager()

	go botServer.run()
	return botServer
}

func (b *BotServer) run() {
	// 运行管理器
	go bot.GetManageBot().Run()

	// 注册模块
	module.RegModules()

	// 运行主循环
	go b.runMainLoop()
}

// runMainLoop 运行主循环
func (b *BotServer) runMainLoop() {
	for count := 0; count < 100000; count++ {
		logger.Warn("main loop start, count: %v", count)
		b.mainLoop()
		logger.Warn("main loop stop, count: %v", count)
	}
}

// mainLoop 主循环
func (b *BotServer) mainLoop() {
	for {
		select {
		case netMsg := <-mq.GetNetMsgChan():
			// 接收gate的消息
			routeManager.HandleNetMsg(netMsg)
		}
	}
}

// Close 关闭服务器
func (b *BotServer) Close() {

}
