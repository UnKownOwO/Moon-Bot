package module

import (
	"moon-bot/bs/bot"
	"moon-bot/bs/module/help"
)

// Modules 注册模块
var Modules = []*bot.ModuleInfo{
	// hello.InitModule(),
	help.InitModule(), // 帮助模块
}
