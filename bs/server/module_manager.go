package server

import (
	"moon-bot/bs/bot"
	"moon-bot/bs/module"
	"moon-bot/pkg/logger"
)

// ModuleManager 模块管理器
type ModuleManager struct {
}

func NewModuleManager() *ModuleManager {
	r := new(ModuleManager)

	// 注册模块
	module.RegisterModules()

	return r
}

// HandleEvent 处理模块事件
func (m *ModuleManager) HandleEvent(bot *bot.Bot, eventId uint16, event bot.ModuleEvent) {
	for _, moduleInfo := range module.Modules {
		eventFunc, ok := moduleInfo.EventRegMap[eventId]
		if !ok {
			logger.Error("module no route, moduleName: %v, eventId: %v", moduleInfo.Name, eventId)
			return
		}
		eventFunc(bot, event)
	}
}
