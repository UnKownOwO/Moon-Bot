package model

import "github.com/golang/protobuf/proto"

type HandlerFunc func(bot *Bot, payloadMsg proto.Message)

type ModuleConfig struct {
	Name          string                 // 模块名
	Version       string                 // 模块版本
	EventRouteMap map[string]HandlerFunc // 事件路由
}

// RegEvent 模块注册事件
func (m *ModuleConfig) RegEvent(cmdName string, handlerFunc HandlerFunc) {
	m.EventRouteMap[cmdName] = handlerFunc
}
