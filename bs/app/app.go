package app

import (
	"moon-bot/bs/bot"
	"moon-bot/common/mq"
	"moon-bot/pkg/logger"
)

func Run() {
	logger.InitLogger()
	logger.Warn("bs start")

	messageQueue := mq.NewMessageQueue(mq.ServerTypeBs)
	defer messageQueue.Close()

	botManager := bot.NewBotManager(messageQueue)
	defer botManager.Close()

	select {}
}
