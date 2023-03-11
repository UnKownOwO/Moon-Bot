package bot

import (
	"github.com/golang/protobuf/proto"
)

type RouteMsg struct {
	UserId     int64
	RouteFunc  RouteFunc
	PayloadMsg proto.Message
}

type RouteFunc func(bot *Bot, payloadMsg proto.Message)
