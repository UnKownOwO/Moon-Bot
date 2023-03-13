package bot

// ModuleEvent 模块事件
type ModuleEvent interface {
	ModuleEvent()
}

// EventFunc 模块事件处理函数
type EventFunc func(bot *Bot, event ModuleEvent)

// EventPriority 事件优先级
type EventPriority uint8

const (
	EventPriorityLowest  EventPriority = iota // 最低
	EventPriorityLow                          // 低
	EventPriorityMiddle                       // 中等
	EventPriorityHigh                         // 高
	EventPriorityHighest                      // 最高
)

// EventRegInfo 模块事件注册信息
type EventRegInfo struct {
	HandleFunc EventFunc     // 事件处理函数
	Priority   EventPriority // 事件优先级
}

// ModuleInfo 模块信息
type ModuleInfo struct {
	Name        string                   // 模块名
	Version     string                   // 模块版本
	EventRegMap map[uint16]*EventRegInfo // 事件路由
}
