package mq

import "github.com/golang/protobuf/proto"

type MsgType uint8

const (
	MsgTypeProto MsgType = iota // proto消息
)

type ProtoMsg struct {
	UserId             int64         `msgpack:"userId"`             // 机器人QQ号
	CmdName            string        `msgpack:"cmdName"`            // 操作名
	PayloadMessage     proto.Message `msgpack:"-"`                  // 将被序列化的消息
	PayloadMessageData []byte        `msgpack:"payloadMessageData"` // 序列化后的消息数据
}

type NetMsg struct {
	MsgType      MsgType   `msgpack:"msgType"`  // 消息类型
	ProtoMsg     *ProtoMsg `msgpack:"protoMsg"` // proto消息
	TargetServer string    `msgpack:"-"`        // 目标服务器
}
