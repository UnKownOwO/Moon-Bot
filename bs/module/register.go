package module

import (
	"moon-bot/bs/bot"
	"moon-bot/bs/module/hello"
)

// Modules 注册模块
var Modules = []*bot.ModuleInfo{
	hello.InitModule(),
}
