package bot

// ModuleEvent 模块事件
type ModuleEvent interface {
	ModuleEvent()
}

// EventFunc 模块事件处理函数
type EventFunc func(bot *Bot, event ModuleEvent)

// ModuleInfo 模块信息
type ModuleInfo struct {
	Name        string               // 模块名
	Version     string               // 模块版本
	EventRegMap map[uint16]EventFunc // 事件路由
}
