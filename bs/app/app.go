package app

import (
	"moon-bot/bs/bot"
	"moon-bot/pkg/logger"
)

func Run() {
	logger.InitLogger()
	logger.Warn("bs start")
	_ = bot.NewBotManager()
	select {}
}
