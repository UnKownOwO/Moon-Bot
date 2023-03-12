package app

import (
	"context"
	"moon-bot/common/mq"
	"moon-bot/gate/net"
	"moon-bot/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context) error {
	logger.InitLogger()
	logger.Warn("gate start")

	mq.InitMessageQueue(mq.ServerTypeGate)
	defer mq.Close()

	connectManager := net.NewConnectManager()
	defer connectManager.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case <-ctx.Done():
			return nil
		case s := <-c:
			logger.Warn("get a signal %s", s.String())
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				logger.Warn("gate exit")
				return nil
			case syscall.SIGHUP:
			default:
				return nil
			}
		}
	}
}
