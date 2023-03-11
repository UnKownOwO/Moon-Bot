package hello

import (
	"moon-bot/bs/bot"
	"moon-bot/bs/event"
	"moon-bot/pkg/logger"
	"moon-bot/protocol/cmd"
	"moon-bot/protocol/pb"
)

type HelloModule struct {
}

func InitModule() *bot.ModuleInfo {
	helloModule := new(HelloModule)
	logger.Debug("hello模块创建")

	return &bot.ModuleInfo{
		Name:    "hello",
		Version: "1.0.0",
		EventRegMap: map[uint16]bot.EventFunc{
			event.ModuleEventIdPrivateMessage: helloModule.privateMessageEvent,
		},
	}
}

func (h *HelloModule) privateMessageEvent(bot *bot.Bot, moduleEvent bot.ModuleEvent) {
	e := moduleEvent.(*event.PrivateMessageEvent)

	logger.Debug("hello模块收到消息通知, moduleEvent: %v", e.Message)

	msg := &pb.SendPrivateMsg{
		UserId: 80520429,
		// GroupId:    0,
		Message:    "nmsl",
		AutoEscape: true,
	}
	bot.SendMsg(cmd.SendPrivateMsg, msg)
}
