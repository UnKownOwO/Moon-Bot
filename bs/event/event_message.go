package event

import "moon-bot/protocol/pb"

// MessageEvent 消息上报事件
type MessageEvent struct {
	SubType    string            // 消息子类型
	MessageId  int32             // 消息 ID
	UserId     int64             // 发送者 QQ 号
	Message    string            // 消息内容
	RawMessage string            // 原始消息内容
	Font       int32             // 字体
	Sender     *pb.MessageSender // 发送人信息
}

// PrivateMessageEvent 私聊消息事件
type PrivateMessageEvent struct {
	*MessageEvent
	TargetId   int64 // 收到私聊消息的机器人 QQ 号
	TempSource int32 // 临时会话来源
}

func (m *PrivateMessageEvent) ModuleEvent() {}

// GroupMessageEvent 群聊消息事件
type GroupMessageEvent struct {
	*MessageEvent
	GroupId   int64                // 群号
	Anonymous *pb.MessageAnonymous // 匿名信息, 如果不是匿名消息则为 null
}

func (m *GroupMessageEvent) ModuleEvent() {}
