package app

import (
	"moon-bot/gate/socket"
	"moon-bot/pkg/logger"
)

func Run() {
	logger.InitLogger()
	logger.Warn("gate start")

	connectManager := socket.NewConnectManager()
	defer connectManager.Close()

	select {}
}
