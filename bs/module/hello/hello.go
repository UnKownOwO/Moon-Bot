package hello

import (
	"moon-bot/bs/event"
	"moon-bot/bs/model"
	"moon-bot/pkg/logger"
)

type HelloModule struct {
}

func InitModule() *model.ModuleInfo {
	helloModule := new(HelloModule)
	logger.Debug("hello模块创建")

	return &model.ModuleInfo{
		Name:    "hello",
		Version: "1.0.0",
		EventRegMap: map[uint16]model.EventFunc{
			event.ModuleEventIdPrivateMessage: helloModule.privateMessageEvent,
		},
	}
}

func (h *HelloModule) privateMessageEvent(bot *model.Bot, moduleEvent model.ModuleEvent) {
	e := moduleEvent.(*event.PrivateMessageEvent)

	logger.Debug("hello模块收到消息通知, moduleEvent: %v", e.Message)
}
