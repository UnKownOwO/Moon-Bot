package bot

import (
	"moon-bot/pkg/logger"
)

// ManageBot Bot管理器
type ManageBot struct {
	botMap       map[int64]*Bot // 创建的所有bot
	connectChan  chan int64     // bot连接管道
	routeMsgChan chan *RouteMsg // 路由消息管道
}

var manageBot *ManageBot

func GetManageBot() *ManageBot {
	if manageBot == nil {
		manageBot = new(ManageBot)
		manageBot.botMap = make(map[int64]*Bot)
		manageBot.connectChan = make(chan int64)
		manageBot.routeMsgChan = make(chan *RouteMsg, 100)
	}
	return manageBot
}

func (m *ManageBot) Run() {
	// 用户连接处理
	for {
		select {
		case userId := <-m.connectChan:
			_, ok := m.botMap[userId]
			if ok {
				logger.Error("bot重连未处理呢")
				continue
			}
			bot := NewBot(userId)
			m.botMap[userId] = bot
		case routeMsg := <-m.routeMsgChan:
			bot, ok := m.botMap[routeMsg.UserId]
			if !ok {
				logger.Error("bot not exist, userId: %v", routeMsg.UserId)
				continue
			}
			bot.GetRouteMsgChan() <- routeMsg
		}
	}
}

// onlineBot bot上线
func (m *ManageBot) onlineBot(userId int64) {
}

// offlineBot bot离线
func (m *ManageBot) offlineBot() {

}

// GetRouteMsgChan 获取路由消息管道
func (m *ManageBot) GetRouteMsgChan() chan<- *RouteMsg {
	return m.routeMsgChan
}

// GetConnectChan 获取用户连接创建用的管道
func (m *ManageBot) GetConnectChan() chan<- int64 {
	return m.connectChan
}
