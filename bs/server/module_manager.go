package server

import (
	"moon-bot/bs/bot"
	"moon-bot/bs/module"
	"moon-bot/pkg/logger"
	"sort"
)

// ModuleManager 模块管理器
type ModuleManager struct {
	eventRegInfoMap map[uint16][]*bot.EventRegInfo // 模块事件注册信息按照优先级排序
}

func NewModuleManager() *ModuleManager {
	r := new(ModuleManager)
	r.eventRegInfoMap = make(map[uint16][]*bot.EventRegInfo)

	// 初始化
	r.initModule()
	return r
}

// initModule 初始化模块
func (m *ModuleManager) initModule() {
	for _, moduleInfo := range module.Modules {
		logger.Info("load module, name: %v, version: %v", moduleInfo.Name, moduleInfo.Version)
		// 存放每个模块的事件注册信息
		for eventId, regInfo := range moduleInfo.EventRegMap {
			// 该事件存放注册信息的列表不存在则创建
			_, ok := m.eventRegInfoMap[eventId]
			if !ok {
				m.eventRegInfoMap[eventId] = make([]*bot.EventRegInfo, 0, 0)
			}
			// 添加事件处理函数
			m.eventRegInfoMap[eventId] = append(m.eventRegInfoMap[eventId], regInfo)
		}
	}
	// 排序模块注册信息优先级 降序
	for _, infoList := range m.eventRegInfoMap {
		sort.Slice(infoList, func(i, j int) bool {
			return infoList[i].Priority < infoList[j].Priority
		})
	}
}

// handleEvent 处理模块事件
func (m *ModuleManager) handleEvent(bot *bot.Bot, eventId uint16, event bot.ModuleEvent) {
	regInfoList, ok := m.eventRegInfoMap[eventId]
	if !ok {
		// logger.Debug("module no route, eventId: %v", eventId)
		return
	}
	// 执行每一个处理函数
	for _, info := range regInfoList {
		info.HandleFunc(bot, event)
	}
}
