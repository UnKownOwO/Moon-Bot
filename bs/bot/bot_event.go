package bot

import (
	"moon-bot/bs/model"
	"moon-bot/pkg/logger"
	"moon-bot/protocol/cmd"

	"github.com/golang/protobuf/proto"
)

func (b *BotManager) messageEvent(bot *model.Bot, payloadMsg proto.Message) {
	logger.Debug("消息事件收到准备通知模块")
	moduleManager.handleEvent(cmd.MessageEvent, bot, payloadMsg)
}
