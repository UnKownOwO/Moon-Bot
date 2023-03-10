package model

type ModuleEvent interface {
	ModuleEvent()
}

type EventFunc func(bot *Bot, event ModuleEvent)

type ModuleInfo struct {
	Name        string               // 模块名
	Version     string               // 模块版本
	EventRegMap map[uint16]EventFunc // 事件路由
}
