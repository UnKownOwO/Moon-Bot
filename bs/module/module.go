package module

import (
	"moon-bot/bs/bot"
	"moon-bot/bs/module/hello"
)

var Modules []*bot.ModuleInfo

// RegModules 注册模块
func RegModules() {
	Modules = []*bot.ModuleInfo{
		hello.InitModule(), // 第一个模块
	}
}
