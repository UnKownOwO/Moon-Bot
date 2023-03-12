package server

import (
	"moon-bot/bs/bot"
	"moon-bot/common/mq"
	"moon-bot/pkg/logger"
)

var (
	routeManager  *RouteManager
	moduleManager *ModuleManager
)

type BotServer struct {
}

func NewBotServer() *BotServer {
	botServer := new(BotServer)

	go botServer.run()
	return botServer
}

func (b *BotServer) run() {
	// 初始化服务器管理器
	moduleManager = NewModuleManager()
	routeManager = NewRouteManager()

	// 运行管理器
	go bot.GetManageBot().Run()

	// 运行主循环
	go b.runMainLoop()
}

// runMainLoop 运行主循环
func (b *BotServer) runMainLoop() {
	for count := 0; count < 100000; count++ {
		logger.Warn("server main loop start, count: %v", count)
		b.mainLoop()
		logger.Warn("server main loop stop, count: %v", count)
	}
}

// mainLoop 主循环
func (b *BotServer) mainLoop() {
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! SERVER MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
		}
	}()
	for {
		select {
		case netMsg := <-mq.GetNetMsgChan():
			// 接收gate的消息
			routeManager.handleNetMsg(netMsg)
		}
	}
}

// Close 关闭服务器
func (b *BotServer) Close() {
	bot.GetManageBot().Close()
}
