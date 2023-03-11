package server

import (
	"moon-bot/bs/bot"
	"moon-bot/bs/module"
	"moon-bot/pkg/logger"
)

// EventManager 事件管理器
type EventManager struct {
}

func NewEventManager() *EventManager {
	r := new(EventManager)

	return r
}

// HandleModuleEvent 处理模块事件
func (m *EventManager) HandleModuleEvent(eventId uint16, bot *bot.Bot, event bot.ModuleEvent) {
	for _, moduleInfo := range module.Modules {
		eventFunc, ok := moduleInfo.EventRegMap[eventId]
		if !ok {
			logger.Error("module no route, moduleName: %v, eventId: %v", moduleInfo.Name, eventId)
			return
		}
		eventFunc(bot, event)
	}
}
