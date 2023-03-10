// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: event/event_meta.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 心跳包的 status 字段的 stat 字段
type MetaStatusStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PacketReceived  uint64 `protobuf:"varint,1,opt,name=packet_received,json=packetReceived,proto3" json:"packet_received,omitempty"`      // 收包数
	PacketSent      uint64 `protobuf:"varint,2,opt,name=packet_sent,json=packetSent,proto3" json:"packet_sent,omitempty"`                  // 发包数
	PacketLost      uint64 `protobuf:"varint,3,opt,name=packet_lost,json=packetLost,proto3" json:"packet_lost,omitempty"`                  // 丢包数
	MessageReceived uint64 `protobuf:"varint,4,opt,name=message_received,json=messageReceived,proto3" json:"message_received,omitempty"`   // 消息接收数
	MessageSent     uint64 `protobuf:"varint,5,opt,name=message_sent,json=messageSent,proto3" json:"message_sent,omitempty"`               // 消息发送数
	DisconnectTimes uint32 `protobuf:"varint,6,opt,name=disconnect_times,json=disconnectTimes,proto3" json:"disconnect_times,omitempty"`   // 连接断开次数
	LostTimes       uint64 `protobuf:"varint,7,opt,name=lost_times,json=lostTimes,proto3" json:"lost_times,omitempty"`                     // 连接丢失次数
	LastMessageTime int64  `protobuf:"varint,8,opt,name=last_message_time,json=lastMessageTime,proto3" json:"last_message_time,omitempty"` // 最后一次消息时间
}

func (x *MetaStatusStatistics) Reset() {
	*x = MetaStatusStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_event_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaStatusStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaStatusStatistics) ProtoMessage() {}

func (x *MetaStatusStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_event_event_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaStatusStatistics.ProtoReflect.Descriptor instead.
func (*MetaStatusStatistics) Descriptor() ([]byte, []int) {
	return file_event_event_meta_proto_rawDescGZIP(), []int{0}
}

func (x *MetaStatusStatistics) GetPacketReceived() uint64 {
	if x != nil {
		return x.PacketReceived
	}
	return 0
}

func (x *MetaStatusStatistics) GetPacketSent() uint64 {
	if x != nil {
		return x.PacketSent
	}
	return 0
}

func (x *MetaStatusStatistics) GetPacketLost() uint64 {
	if x != nil {
		return x.PacketLost
	}
	return 0
}

func (x *MetaStatusStatistics) GetMessageReceived() uint64 {
	if x != nil {
		return x.MessageReceived
	}
	return 0
}

func (x *MetaStatusStatistics) GetMessageSent() uint64 {
	if x != nil {
		return x.MessageSent
	}
	return 0
}

func (x *MetaStatusStatistics) GetDisconnectTimes() uint32 {
	if x != nil {
		return x.DisconnectTimes
	}
	return 0
}

func (x *MetaStatusStatistics) GetLostTimes() uint64 {
	if x != nil {
		return x.LostTimes
	}
	return 0
}

func (x *MetaStatusStatistics) GetLastMessageTime() int64 {
	if x != nil {
		return x.LastMessageTime
	}
	return 0
}

// 心跳包上报中作为成员使用
type MetaStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppInitialized bool                  `protobuf:"varint,1,opt,name=app_initialized,json=appInitialized,proto3" json:"app_initialized,omitempty"` // 程序是否初始化完毕
	AppEnabled     bool                  `protobuf:"varint,2,opt,name=app_enabled,json=appEnabled,proto3" json:"app_enabled,omitempty"`             // 程序是否可用
	PluginsGood    bool                  `protobuf:"varint,3,opt,name=plugins_good,json=pluginsGood,proto3" json:"plugins_good,omitempty"`          // 插件正常(可能为 null)
	AppGood        bool                  `protobuf:"varint,4,opt,name=app_good,json=appGood,proto3" json:"app_good,omitempty"`                      // 程序正常
	Good           bool                  `protobuf:"varint,5,opt,name=good,proto3" json:"good,omitempty"`                                           // 运行正常
	Online         bool                  `protobuf:"varint,6,opt,name=online,proto3" json:"online,omitempty"`                                       // 是否在线
	Stat           *MetaStatusStatistics `protobuf:"bytes,7,opt,name=stat,proto3" json:"stat,omitempty"`                                            // 统计信息
}

func (x *MetaStatus) Reset() {
	*x = MetaStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_event_meta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaStatus) ProtoMessage() {}

func (x *MetaStatus) ProtoReflect() protoreflect.Message {
	mi := &file_event_event_meta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaStatus.ProtoReflect.Descriptor instead.
func (*MetaStatus) Descriptor() ([]byte, []int) {
	return file_event_event_meta_proto_rawDescGZIP(), []int{1}
}

func (x *MetaStatus) GetAppInitialized() bool {
	if x != nil {
		return x.AppInitialized
	}
	return false
}

func (x *MetaStatus) GetAppEnabled() bool {
	if x != nil {
		return x.AppEnabled
	}
	return false
}

func (x *MetaStatus) GetPluginsGood() bool {
	if x != nil {
		return x.PluginsGood
	}
	return false
}

func (x *MetaStatus) GetAppGood() bool {
	if x != nil {
		return x.AppGood
	}
	return false
}

func (x *MetaStatus) GetGood() bool {
	if x != nil {
		return x.Good
	}
	return false
}

func (x *MetaStatus) GetOnline() bool {
	if x != nil {
		return x.Online
	}
	return false
}

func (x *MetaStatus) GetStat() *MetaStatusStatistics {
	if x != nil {
		return x.Stat
	}
	return nil
}

// 元事件上报
type MetaEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time     int64  `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`                        // 事件发生的时间戳
	SelfId   int64  `protobuf:"varint,2,opt,name=self_id,json=selfId,proto3" json:"self_id,omitempty"`      // 收到事件的机器人 QQ 号
	PostType string `protobuf:"bytes,3,opt,name=post_type,json=postType,proto3" json:"post_type,omitempty"` // 上报类型
	// 心跳包
	MetaEventType string      `protobuf:"bytes,4,opt,name=meta_event_type,json=metaEventType,proto3" json:"meta_event_type,omitempty"` // 上报类型
	Status        *MetaStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`                                      // 应用程序状态
	Interval      int64       `protobuf:"varint,6,opt,name=interval,proto3" json:"interval,omitempty"`                                 // 距离上一次心跳包的时间(单位是毫秒)
	// 生命周期
	SubType string `protobuf:"bytes,7,opt,name=sub_type,json=subType,proto3" json:"sub_type,omitempty"` // 子类型
}

func (x *MetaEvent) Reset() {
	*x = MetaEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_event_meta_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaEvent) ProtoMessage() {}

func (x *MetaEvent) ProtoReflect() protoreflect.Message {
	mi := &file_event_event_meta_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaEvent.ProtoReflect.Descriptor instead.
func (*MetaEvent) Descriptor() ([]byte, []int) {
	return file_event_event_meta_proto_rawDescGZIP(), []int{2}
}

func (x *MetaEvent) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *MetaEvent) GetSelfId() int64 {
	if x != nil {
		return x.SelfId
	}
	return 0
}

func (x *MetaEvent) GetPostType() string {
	if x != nil {
		return x.PostType
	}
	return ""
}

func (x *MetaEvent) GetMetaEventType() string {
	if x != nil {
		return x.MetaEventType
	}
	return ""
}

func (x *MetaEvent) GetStatus() *MetaStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *MetaEvent) GetInterval() int64 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *MetaEvent) GetSubType() string {
	if x != nil {
		return x.SubType
	}
	return ""
}

var File_event_event_meta_proto protoreflect.FileDescriptor

var file_event_event_meta_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xc5, 0x02, 0x0a,
	0x14, 0x4d, 0x65, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6c, 0x6f, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x73, 0x74,
	0x12, 0x29, 0x0a, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x29,
	0x0a, 0x10, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x73,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6c,
	0x6f, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0xee, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x61, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x70,
	0x70, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x70, 0x70, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x5f, 0x67, 0x6f, 0x6f, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x47, 0x6f, 0x6f, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x67, 0x6f, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x61, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x67,
	0x6f, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x67, 0x6f, 0x6f, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52,
	0x04, 0x73, 0x74, 0x61, 0x74, 0x22, 0xdc, 0x01, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x6c, 0x66, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x6c, 0x66, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a,
	0x0f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x75, 0x62,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_event_meta_proto_rawDescOnce sync.Once
	file_event_event_meta_proto_rawDescData = file_event_event_meta_proto_rawDesc
)

func file_event_event_meta_proto_rawDescGZIP() []byte {
	file_event_event_meta_proto_rawDescOnce.Do(func() {
		file_event_event_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_event_meta_proto_rawDescData)
	})
	return file_event_event_meta_proto_rawDescData
}

var file_event_event_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_event_event_meta_proto_goTypes = []interface{}{
	(*MetaStatusStatistics)(nil), // 0: pb.MetaStatusStatistics
	(*MetaStatus)(nil),           // 1: pb.MetaStatus
	(*MetaEvent)(nil),            // 2: pb.MetaEvent
}
var file_event_event_meta_proto_depIdxs = []int32{
	0, // 0: pb.MetaStatus.stat:type_name -> pb.MetaStatusStatistics
	1, // 1: pb.MetaEvent.status:type_name -> pb.MetaStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_event_event_meta_proto_init() }
func file_event_event_meta_proto_init() {
	if File_event_event_meta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_event_meta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaStatusStatistics); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_event_event_meta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_event_event_meta_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_event_meta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_event_meta_proto_goTypes,
		DependencyIndexes: file_event_event_meta_proto_depIdxs,
		MessageInfos:      file_event_event_meta_proto_msgTypes,
	}.Build()
	File_event_event_meta_proto = out.File
	file_event_event_meta_proto_rawDesc = nil
	file_event_event_meta_proto_goTypes = nil
	file_event_event_meta_proto_depIdxs = nil
}
