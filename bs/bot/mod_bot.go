package bot

import (
	"moon-bot/common/mq"
	"moon-bot/pkg/logger"

	"github.com/golang/protobuf/proto"
)

// ModBot 基础模块
type ModBot struct {
	routeMsgChan chan *RouteMsg // 路由消息管道
	closeChan    chan bool      // 关闭bot管道
	bot          *Bot
}

// initData 初始化模块
func (m *ModBot) initData() {
	m.routeMsgChan = make(chan *RouteMsg, 100)
	m.closeChan = make(chan bool, 1)
}

// loadData 加载数据
func (m *ModBot) loadData(bot *Bot) {
	m.bot = bot
	m.initData()

	// 运行bot主循环
	go m.botMainLoop()
}

// saveData 保存数据
func (m *ModBot) saveData() {
}

// botMainLoop 路由消息主循环
func (m *ModBot) botMainLoop() {
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! BOT MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
			logger.Error("the motherfucker bot userId: %v", m.bot.UserId)
			GetManageBot().kickBot(m.bot.UserId)
		}
	}()
	for {
		select {
		case routeMsg := <-m.routeMsgChan:
			// 处理路由消息
			routeMsg.RouteFunc(m.bot, routeMsg.PayloadMsg)
		case <-m.closeChan:
			// 关闭bot
			return
		}
	}
}

// 对外接口

// SendMsg 发送消息给客户端
func (b *Bot) SendMsg(cmdName string, payloadMsg proto.Message) {
	// 发送消息到gate
	mq.Send(mq.ServerTypeGate, &mq.NetMsg{
		MsgType: mq.MsgTypeProto,
		ProtoMsg: &mq.ProtoMsg{
			UserId:         b.UserId,
			CmdName:        cmdName,
			PayloadMessage: payloadMsg,
		},
	})
}
