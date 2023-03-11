package app

import (
	"moon-bot/bs/server"
	"moon-bot/common/mq"
	"moon-bot/pkg/logger"
)

func Run() {
	logger.InitLogger()
	logger.Warn("bs start")

	mq.InitMessageQueue(mq.ServerTypeBs)
	defer mq.Close()

	botServer := server.NewBotServer()
	defer botServer.Close()

	select {}
}
