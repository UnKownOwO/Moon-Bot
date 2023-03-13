package help

import (
	"moon-bot/bs/bot"
	"moon-bot/bs/event"
	"moon-bot/protocol/cmd"
	"moon-bot/protocol/pb"
)

// InitModule 初始化模块
func InitModule() *bot.ModuleInfo {
	return &bot.ModuleInfo{
		Name:    "帮助",
		Version: "1.0.0",
		EventRegMap: map[uint16]*bot.EventRegInfo{
			event.ModuleEventIdGroupMessage: {
				HandleFunc: GroupMessageEvent,
				Priority:   bot.EventPriorityMiddle,
			},
		},
	}
}

// GroupMessageEvent 群消息
func GroupMessageEvent(bot *bot.Bot, moduleEvent bot.ModuleEvent) {
	e := moduleEvent.(*event.GroupMessageEvent)

	switch e.Message {
	case "help", "菜单", "功能", "帮助":
		msg := &pb.SendGroupMsg{
			GroupId:    e.GroupId,
			Message:    "这是一个菜单命令",
			AutoEscape: false,
		}
		bot.SendMsg(cmd.SendGroupMsg, msg)
	}
}
