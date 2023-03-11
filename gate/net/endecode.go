package net

import (
	"encoding/json"
	"moon-bot/pkg/logger"
	"moon-bot/protocol/cmd"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

type ProtoMessage struct {
	CmdName        string
	PayloadMessage proto.Message
}

// ApiJson websocket调用api的json结构
type ApiJson struct {
	Action string        `json:"action"` // 终结点名称
	Params proto.Message `json:"params"` // 参数
	Echo   string        `json:"echo"`   // 回声
}

// EventJson 上报事件json结构
type EventJson struct {
	Time     int64  `json:"time"`      // 事件发生的unix时间戳
	SelfId   int64  `json:"self_id"`   // 收到事件的机器人的 QQ 号
	PostType string `json:"post_type"` // 表示该上报的类型
}

// protoMessageToApiJson protoMessage转换为api的json格式
func (c *ConnectManager) protoMessageToApiJson(protoMessage *ProtoMessage) []byte {
	apiJson := &ApiJson{
		Action: protoMessage.CmdName,
		Params: protoMessage.PayloadMessage,
		Echo:   "",
	}
	jsonData, err := json.Marshal(apiJson)
	if err != nil {
		logger.Error("marshal json to data error: %v, apiJson: %v", err, apiJson)
		return nil
	}
	return jsonData
}

// eventJsonToProtoMessage 上报事件json格式转换为protoMessage
func (c *ConnectManager) eventJsonToProtoMessage(data []byte) *ProtoMessage {
	var eventJson EventJson
	err := json.Unmarshal(data, &eventJson)
	if err != nil {
		logger.Error("unmarshal event json error: %v", err)
		return nil
	}
	// 获取消息对象
	protoObj := cmd.GetCmdObjManager().GetProtoObjByCmdName(eventJson.PostType)
	if protoObj == nil {
		logger.Error("get new proto object is nil, data: %v", string(data))
		return nil
	}
	// 将json的数据转换到proto
	err = jsonpb.UnmarshalString(string(data), protoObj)
	if err != nil {
		logger.Error("unmarshal json to proto error: %v", err)
		return nil
	}
	protoMessage := &ProtoMessage{
		CmdName:        eventJson.PostType,
		PayloadMessage: protoObj,
	}
	return protoMessage
}
