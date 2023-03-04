package hello

import (
	"moon-bot/bs/model"
	"moon-bot/pkg/logger"
	"moon-bot/protocol/cmd"

	"github.com/golang/protobuf/proto"
)

type HelloModule struct {
}

func NewHelloModule() *HelloModule {
	helloModule := new(HelloModule)
	logger.Debug("hello模块创建")
	return helloModule
}

func (h *HelloModule) InitModule(config *model.ModuleConfig) {
	config.RegEvent(cmd.MessageEvent, h.messageEvent)
	logger.Debug("hello模块初始化")
}

func (h *HelloModule) messageEvent(bot *model.Bot, payloadMsg proto.Message) {
	logger.Debug("hello模块收到消息通知, payloadMsg: %v", payloadMsg)
}
