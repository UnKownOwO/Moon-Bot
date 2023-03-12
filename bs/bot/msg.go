package bot

import (
	"github.com/golang/protobuf/proto"
)

// UserCtrlType 用户控制类型
type UserCtrlType uint8

const (
	UserCtrlTypeConnect     UserCtrlType = iota // 用户上线
	UserCtrlTypeDisconnect                      // 用户离线
	UserCtrlTypeServerClose                     // 服务器关闭
)

// UserCtrlMsg 用户管理消息
type UserCtrlMsg struct {
	UserCtrlType UserCtrlType // 用户控制类型
	UserId       int64        // 用户Id
}

// RouteFunc 路由处理函数
type RouteFunc func(bot *Bot, payloadMsg proto.Message)

// RouteMsg 路由消息
type RouteMsg struct {
	UserId     int64         // 用户Id
	CmdName    string        // 操作名
	RouteFunc  RouteFunc     // 路由处理函数
	PayloadMsg proto.Message // proto消息
}
