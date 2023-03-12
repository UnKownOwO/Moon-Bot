package app

import (
	"context"
	"moon-bot/bs/server"
	"moon-bot/common/mq"
	"moon-bot/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context) error {
	logger.InitLogger()
	logger.Warn("bs start")

	mq.InitMessageQueue(mq.ServerTypeBs)
	defer mq.Close()

	botServer := server.NewBotServer()
	defer botServer.Close()

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
				logger.Warn("bs exit")
				return nil
			case syscall.SIGHUP:
			default:
				return nil
			}
		}
	}
}
