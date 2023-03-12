package bot

import (
	"moon-bot/common/mq"
	"moon-bot/pkg/logger"
	"time"
)

const (
	CheckOffBotTime = time.Minute * 30 // 检测离线bot并清理时间
)

// ManageBot Bot管理器
type ManageBot struct {
	botMap          map[int64]*Bot    // 创建的所有bot
	userCtrlMsgChan chan *UserCtrlMsg // 用户管理消息管道
	routeMsgChan    chan *RouteMsg    // 路由消息管道
}

var manageBot *ManageBot

func GetManageBot() *ManageBot {
	if manageBot == nil {
		manageBot = new(ManageBot)
		manageBot.botMap = make(map[int64]*Bot)
		manageBot.userCtrlMsgChan = make(chan *UserCtrlMsg)
		manageBot.routeMsgChan = make(chan *RouteMsg, 100)
	}
	return manageBot
}

func (m *ManageBot) Run() {
	// 清理离线bot定时器
	// 通常生产环境下30分钟清理一次
	// 内存换性能 一般bot也不会太多
	// 每次掉线重连都得创建太浪费性能
	checkOffTicker := time.NewTicker(CheckOffBotTime)
	for {
		select {
		case userCtrlMsg := <-m.userCtrlMsgChan:
			// 用户管理消息
			switch userCtrlMsg.UserCtrlType {
			case UserCtrlTypeConnect:
				// 用户连接
				m.onlineBot(userCtrlMsg.UserId)
			case UserCtrlTypeDisconnect:
				// 用户离线
				m.offlineBot(userCtrlMsg.UserId)
			case UserCtrlTypeServerClose:
				// 服务器关闭
				for _, bot := range m.botMap {
					// 踢出所有用户
					m.kickBot(bot.UserId)
				}
			}
		case routeMsg := <-m.routeMsgChan:
			// 用户路由消息
			bot, ok := m.botMap[routeMsg.UserId]
			if !ok {
				logger.Error("bot not exist, userId: %v", routeMsg.UserId)
				continue
			}
			bot.GetRouteMsgChan() <- routeMsg
		case <-checkOffTicker.C:
			// 清理离线bot
			m.checkOffBot()
		}
	}
}

func (m *ManageBot) Close() {
	// 通知服务器即将关闭
	m.userCtrlMsgChan <- &UserCtrlMsg{
		UserCtrlType: UserCtrlTypeServerClose,
	}
}

// onlineBot bot上线
func (m *ManageBot) onlineBot(userId int64) {
	bot, ok := m.botMap[userId]
	if ok {
		// 重连处理
		if bot.exitTime != 0 {
			logger.Info("bot reconnect, userId: %v", userId)
			bot.exitTime = 0
		}
		return
	}
	// 没有存储的bot数据则创建
	logger.Info("bot connect, userId: %v", userId)
	bot = NewBot(userId)
	m.botMap[userId] = bot
}

// offlineBot bot离线
func (m *ManageBot) offlineBot(userId int64) {
	bot, ok := m.botMap[userId]
	if !ok {
		logger.Error("bot not exist, userId: %v", userId)
		return
	}
	// 用户是否已离线
	if bot.exitTime != 0 {
		logger.Error("bot has been offline, userId: %v", userId)
		return
	}
	logger.Info("bot disconnect, userId: %v", userId)
	bot.exitTime = time.Now().Unix()
}

// kickBot 踢出bot
func (m *ManageBot) kickBot(userId int64) {
	// 发送消息到gate关闭连接
	mq.Send(mq.ServerTypeGate, &mq.NetMsg{
		MsgType: mq.MsgTypeOffline,
		OfflineMsg: &mq.OfflineMsg{
			UserId: userId,
		},
	})
}

// checkOffBot 检查离线的bot并清理
func (m *ManageBot) checkOffBot() {
	logger.Debug("check offline bot start")
	for _, bot := range m.botMap {
		// 离线时间等于0代表未离线
		if bot.exitTime != 0 && bot.exitTime < time.Now().Unix() {
			logger.Debug("bot clean, userId: %v, exitTime: %v", bot.UserId, bot.exitTime)
			bot.close()
			delete(m.botMap, bot.UserId)
		}
	}
	logger.Debug("check offline bot finish")
}

// GetRouteMsgChan 获取路由消息管道
func (m *ManageBot) GetRouteMsgChan() chan<- *RouteMsg {
	return m.routeMsgChan
}

// GetUserCtrlMsgChan 获取用户控制消息管道
func (m *ManageBot) GetUserCtrlMsgChan() chan<- *UserCtrlMsg {
	return m.userCtrlMsgChan
}
