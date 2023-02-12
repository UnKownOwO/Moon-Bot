package app

import (
	"moon-bot/common/mq"
	"moon-bot/gate/net"
	"moon-bot/pkg/logger"
)

func Run() {
	logger.InitLogger()
	logger.Warn("gate start")

	messageQueue := mq.NewMessageQueue(mq.ServerTypeGate)
	defer messageQueue.Close()

	connectManager := net.NewConnectManager(messageQueue)
	defer connectManager.Close()

	select {}
}
