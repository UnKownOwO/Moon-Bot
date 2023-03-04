package bot

import (
	"moon-bot/bs/model"
	"moon-bot/bs/module/hello"
	"moon-bot/pkg/logger"

	"github.com/golang/protobuf/proto"
)

type ModuleI interface {
	InitModule(*model.ModuleConfig) // 初始化模块
}

type ModuleManager struct {
	moduleMap map[string]*model.ModuleConfig
}

func NewModuleManager() *ModuleManager {
	r := new(ModuleManager)
	r.moduleMap = make(map[string]*model.ModuleConfig)

	r.regAllModule()
	return r
}

// regAllModule 注册所有模块
func (m *ModuleManager) regAllModule() {
	m.regModule("hello", "1.0.0", hello.NewHelloModule())
}

// handleEvent 处理模块事件
func (m *ModuleManager) handleEvent(cmdName string, bot *model.Bot, payloadMsg proto.Message) {
	for _, config := range m.moduleMap {
		handlerFunc, ok := config.EventRouteMap[cmdName]
		if !ok {
			logger.Error("module no route, moduleName: %v, cmdName: %v", config.Name, cmdName)
			return
		}
		handlerFunc(bot, payloadMsg)
	}
}

// RegModule 注册模块
func (m *ModuleManager) regModule(name string, version string, module ModuleI) {
	config := &model.ModuleConfig{
		Name:          name,
		Version:       version,
		EventRouteMap: make(map[string]model.HandlerFunc),
	}
	// 空校验
	if config == nil {
		logger.Error("config config is nil, name: %v, version: %v", name, version)
		return
	}
	// 重复校验
	_, ok := m.moduleMap[name]
	if ok {
		logger.Error("config has both, name: %v, version: %v", name, version)
		return
	}
	// 初始化模块
	module.InitModule(config)
	// 记录模块
	m.moduleMap[config.Name] = config
}
