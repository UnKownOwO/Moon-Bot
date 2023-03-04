package bot

import (
	"moon-bot/common/mq"
	"moon-bot/pkg/logger"
)

var (
	botManager    *BotManager
	routeManager  *RouteManager
	loginManager  *LoginManager
	moduleManager *ModuleManager
)

type BotManager struct {
	messageQueue *mq.MessageQueue
}

func NewBotManager(messageQueue *mq.MessageQueue) *BotManager {
	r := new(BotManager)
	r.messageQueue = messageQueue

	// 初始化管理器
	botManager = r
	routeManager = NewRouteManager()
	loginManager = NewLoginManager()
	moduleManager = NewModuleManager()

	go botManager.runMainLoop()
	return botManager
}

// runMainLoop 运行主循环
func (b *BotManager) runMainLoop() {
	for count := 0; count < 100000; count++ {
		logger.Warn("main loop start, count: %v", count)
		b.mainLoop()
		logger.Warn("main loop stop, count: %v", count)
	}
}

// mainLoop 主循环
func (b *BotManager) mainLoop() {
	for {
		select {
		case netMsg := <-b.messageQueue.GetNetMsgChan():
			// 接收gate消息
			routeManager.HandleNetMsg(netMsg)
		}
	}
}

// Close 服务器关闭
func (b *BotManager) Close() {

}
