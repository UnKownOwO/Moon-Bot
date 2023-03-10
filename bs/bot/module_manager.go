package bot

import (
	"moon-bot/bs/model"
	"moon-bot/bs/module/hello"
	"moon-bot/pkg/logger"
)

type ModuleManager struct {
	moduleMap map[string]*model.ModuleInfo
}

func NewModuleManager() *ModuleManager {
	r := new(ModuleManager)
	r.moduleMap = make(map[string]*model.ModuleInfo)

	r.regAllModule()
	return r
}

// regAllModule 注册所有模块
func (m *ModuleManager) regAllModule() {
	m.regModule([]*model.ModuleInfo{
		hello.InitModule(), // 第一个模块
	}...)
}

// RegModule 注册模块
func (m *ModuleManager) regModule(moduleInfoList ...*model.ModuleInfo) {
	for _, info := range moduleInfoList {
		// 重复校验
		_, ok := m.moduleMap[info.Name]
		if ok {
			logger.Error("module has both, moduleInfo: %v", info)
			return
		}
		// 记录模块
		m.moduleMap[info.Name] = info
	}
}

// handleEvent 处理模块事件
func (m *ModuleManager) handleEvent(eventId uint16, bot *model.Bot, event model.ModuleEvent) {
	for _, config := range m.moduleMap {
		eventFunc, ok := config.EventRegMap[eventId]
		if !ok {
			logger.Error("module no route, moduleName: %v, eventId: %v", config.Name, eventId)
			return
		}
		eventFunc(bot, event)
	}
}
