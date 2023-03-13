package server

import (
	"moon-bot/bs/bot"
	"moon-bot/common/constant"
	"moon-bot/common/mq"
	"moon-bot/pkg/logger"
	"moon-bot/protocol/cmd"
	"moon-bot/protocol/pb"
	"time"
)

var (
	routeManager  *RouteManager
	moduleManager *ModuleManager
)

type BotServer struct {
}

func NewBotServer() *BotServer {
	botServer := new(BotServer)

	botServer.run()
	return botServer
}

func (b *BotServer) run() {
	// 初始化服务器管理器
	moduleManager = NewModuleManager()
	routeManager = NewRouteManager()

	// 运行管理器
	go bot.GetManageBot().Run()

	// 运行主循环
	go b.runMainLoop()
}

// runMainLoop 运行主循环
func (b *BotServer) runMainLoop() {
	for count := 0; count < 100000; count++ {
		logger.Warn("server main loop start, count: %v", count)
		b.mainLoop()
		logger.Warn("server main loop stop, count: %v", count)
	}
}

// mainLoop 主循环
func (b *BotServer) mainLoop() {
	// panic捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Error("!!! SERVER MAIN LOOP PANIC !!!")
			logger.Error("error: %v", err)
			logger.Error("stack: %v", logger.Stack())
		}
	}()
	// 通知gate bs已启动重传的定时器
	serverMsgTicker := time.NewTicker(time.Nanosecond)
	for {
		select {
		case netMsg := <-mq.GetNetMsgChan():
			// 接收gate的消息
			switch netMsg.MsgType {
			case mq.MsgTypeProto:
				protoMsg := netMsg.ProtoMsg
				// 确保消息为生命周期连接
				if protoMsg.CmdName == cmd.MetaEvent && protoMsg.PayloadMessage.(*pb.MetaEvent).SubType == constant.MetaEventTypeLifecycleSubTypeConnect {
					// 用户连接
					bot.GetManageBot().GetUserCtrlMsgChan() <- &bot.UserCtrlMsg{
						UserCtrlType: bot.UserCtrlTypeConnect,
						UserId:       protoMsg.UserId,
					}
					continue
				}
				routeManager.doRoute(protoMsg.CmdName, protoMsg.UserId, protoMsg.PayloadMessage)
			case mq.MsgTypeOffline:
				// 用户离线消息
				offlineMsg := netMsg.OfflineMsg
				bot.GetManageBot().GetUserCtrlMsgChan() <- &bot.UserCtrlMsg{
					UserCtrlType: bot.UserCtrlTypeDisconnect,
					UserId:       offlineMsg.UserId,
				}
			case mq.MsgTypeServer:
				// 服务器消息
				serverMsg := netMsg.ServerMsg
				// 确保是gate的消息
				if serverMsg.TargetServer != mq.ServerTypeGate {
					continue
				}
				switch serverMsg.ActionType {
				case mq.ServerMsgActionTypeStart:
					logger.Warn("gate start, start ticker")
					// 重新开始重传
					serverMsgTicker.Reset(time.Nanosecond)
				case mq.ServerMsgActionTypeStartAck:
					logger.Warn("gate ack, stop ticker")
					// 停止重传
					serverMsgTicker.Stop()
				case mq.ServerMsgActionTypeExit:
					logger.Warn("gate exit, close all bot")
					bot.GetManageBot().GetUserCtrlMsgChan() <- &bot.UserCtrlMsg{
						UserCtrlType: bot.UserCtrlTypeServerClose,
					}
				}
			}
		case <-serverMsgTicker.C:
			// 第一次发送为立即发送
			// 发送完后每10s重传一次
			serverMsgTicker.Reset(time.Second * 10)
			// 通知gate bs已启动
			mq.Send(mq.ServerTypeGate, &mq.NetMsg{
				MsgType: mq.MsgTypeServer,
				ServerMsg: &mq.ServerMsg{
					ActionType:   mq.ServerMsgActionTypeStart,
					TargetServer: mq.ServerTypeBs,
				},
			})
		}
	}
}

// Close 关闭服务器
func (b *BotServer) Close() {
	bot.GetManageBot().Close()
	// 通知gate bs关闭
	mq.Send(mq.ServerTypeGate, &mq.NetMsg{
		MsgType: mq.MsgTypeServer,
		ServerMsg: &mq.ServerMsg{
			ActionType:   mq.ServerMsgActionTypeExit,
			TargetServer: mq.ServerTypeBs,
		},
	})
}
