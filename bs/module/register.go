package module

import (
	"moon-bot/bs/bot"
	"moon-bot/bs/module/hello"
)

var Modules []*bot.ModuleInfo

// RegisterModules 注册所有模块
func RegisterModules() {
	Modules = []*bot.ModuleInfo{
		hello.InitModule(),
	}
}
