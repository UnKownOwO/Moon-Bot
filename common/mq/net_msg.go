package mq

import "github.com/golang/protobuf/proto"

type MsgType uint8

const (
	MsgTypeProto   MsgType = iota // proto消息
	MsgTypeOffline                // 用户下线消息
)

type OfflineMsg struct {
	UserId int64 `msgpack:"userId"` // 用户Id
}

type ProtoMsg struct {
	UserId             int64         `msgpack:"userId"`             // 用户Id
	CmdName            string        `msgpack:"cmdName"`            // 操作名
	PayloadMessage     proto.Message `msgpack:"-"`                  // 将被序列化的消息
	PayloadMessageData []byte        `msgpack:"payloadMessageData"` // 序列化后的消息数据
}

type NetMsg struct {
	MsgType      MsgType     `msgpack:"msgType"`    // 消息类型
	ProtoMsg     *ProtoMsg   `msgpack:"protoMsg"`   // proto消息
	OfflineMsg   *OfflineMsg `msgpack:"offlineMsg"` // 用户下线消息
	TargetServer ServerType  `msgpack:"-"`          // 目标服务器
}
