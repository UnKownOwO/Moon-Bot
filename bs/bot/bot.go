package bot

import (
	"moon-bot/common/mq"

	"github.com/golang/protobuf/proto"
)

type ModBase interface {
	InitMod() // 初始化模块
}

const (
	ModIdBot uint8 = iota
)

type Bot struct {
	UserId       int64             // 用户Id
	modManage    map[uint8]ModBase // 模块
	routeMsgChan chan *RouteMsg    // 路由消息管道
}

func NewBot(userId int64) *Bot {
	bot := new(Bot)
	bot.UserId = userId
	bot.modManage = map[uint8]ModBase{
		ModIdBot: new(ModBot),
	}
	bot.routeMsgChan = make(chan *RouteMsg, 100)
	// 初始化模块
	bot.initMod()

	// 路由消息处理
	go bot.routeMsgHandler()
	return bot
}

// routeMsgHandler 路由消息处理
func (b *Bot) routeMsgHandler() {
	for {
		routeMsg := <-b.routeMsgChan
		routeMsg.RouteFunc(b, routeMsg.PayloadMsg)
	}
}

// initMod 初始化模块
func (b *Bot) initMod() {
	for _, mod := range b.modManage {
		mod.InitMod()
	}
}

// SendMsg 发送消息给客户端
func (b *Bot) SendMsg(cmdName string, payloadMsg proto.Message) {
	// TODO 判断机器人是否存在
	// 发送消息到gate
	protoMsg := new(mq.ProtoMsg)
	protoMsg.UserId = b.UserId
	protoMsg.CmdName = cmdName
	protoMsg.PayloadMessage = payloadMsg
	mq.Send(mq.ServerTypeGate, &mq.NetMsg{
		MsgType:  mq.MsgTypeProto,
		ProtoMsg: protoMsg,
	})
}

// GetRouteMsgChan 获取路由消息管道
func (b *Bot) GetRouteMsgChan() chan<- *RouteMsg {
	return b.routeMsgChan
}

// GetMod 获取某一模块
func (b *Bot) GetMod(modId uint8) ModBase {
	return b.modManage[modId]
}

// GetModBot 获取基础模块
func (b *Bot) GetModBot() *ModBot {
	return b.GetMod(ModIdBot).(*ModBot)
}
