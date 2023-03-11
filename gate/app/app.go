package app

import (
	"moon-bot/common/mq"
	"moon-bot/gate/net"
	"moon-bot/pkg/logger"
)

func Run() {
	logger.InitLogger()
	logger.Warn("gate start")

	mq.InitMessageQueue(mq.ServerTypeGate)
	defer mq.Close()

	connectManager := net.NewConnectManager()
	defer connectManager.Close()

	select {}
}
