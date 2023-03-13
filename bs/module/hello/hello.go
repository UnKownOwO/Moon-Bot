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

	return &bot.ModuleInfo{
		Name:    "hello",
		Version: "1.0.0",
		EventRegMap: map[uint16]*bot.EventRegInfo{
			event.ModuleEventIdPrivateMessage: {
				HandleFunc: helloModule.PrivateMessageEvent,
				Priority:   bot.EventPriorityMiddle,
			},
			event.ModuleEventIdGroupMessage: {
				HandleFunc: helloModule.GroupMessageEvent,
				Priority:   bot.EventPriorityLowest,
			},
		},
	}
}

func (h *HelloModule) PrivateMessageEvent(bot *bot.Bot, moduleEvent bot.ModuleEvent) {
	e := moduleEvent.(*event.PrivateMessageEvent)
	logger.Debug("私聊信息: %v", e.Message)

	msg := &pb.SendPrivateMsg{
		UserId: e.UserId,
		// GroupId:    0,
		Message:    "nmsl",
		AutoEscape: false,
	}
	bot.SendMsg(cmd.SendPrivateMsg, msg)
}

func (h *HelloModule) GroupMessageEvent(bot *bot.Bot, moduleEvent bot.ModuleEvent) {
	e := moduleEvent.(*event.GroupMessageEvent)
	logger.Debug("群聊信息: %v", e.Message)

	msg := &pb.SendGroupMsg{
		GroupId:    e.GroupId,
		Message:    "hello~ test!! [CQ:face,id=123][CQ:face,id=123]",
		AutoEscape: false,
	}
	bot.SendMsg(cmd.SendGroupMsg, msg)
}
