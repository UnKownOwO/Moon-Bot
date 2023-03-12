package cmd

import (
	"moon-bot/pkg/logger"
	"moon-bot/protocol/pb"
	"reflect"

	"github.com/golang/protobuf/proto"
)

type CmdObjManager struct {
	cmdNameProtoObjMap map[string]reflect.Type
	protoObjCmdNameMap map[reflect.Type]string
}

var cmdObjManager *CmdObjManager

func GetCmdObjManager() *CmdObjManager {
	if cmdObjManager == nil {
		cmdObjManager = new(CmdObjManager)
		cmdObjManager.cmdNameProtoObjMap = make(map[string]reflect.Type)
		cmdObjManager.protoObjCmdNameMap = make(map[reflect.Type]string)
		cmdObjManager.initAllMessage()
	}
	return cmdObjManager
}

// initMessage 初始化所有消息
func (c *CmdObjManager) initAllMessage() {
	{
		// 事件上报
		c.registerMessage(MessageEvent, &pb.MessageEvent{})
		c.registerMessage(MetaEvent, &pb.MetaEvent{})
	}
	{
		// 请求API
		c.registerMessage(SendPrivateMsg, &pb.SendPrivateMsg{})
		c.registerMessage(SendGroupMsg, &pb.SendGroupMsg{})
	}
}

// registerMessage 注册消息
func (c *CmdObjManager) registerMessage(cmdName string, protoObj proto.Message) {
	// 检查消息是否被重复注册
	_, ok := c.cmdNameProtoObjMap[cmdName]
	if ok {
		logger.Error("register message repeat, cmdName: %v", cmdName)
		return
	}
	refType := reflect.TypeOf(protoObj)
	// 注册消息的反射类型
	c.cmdNameProtoObjMap[cmdName] = refType
	c.protoObjCmdNameMap[refType] = cmdName
}

// GetProtoObjByCmdName 通过操作名获取消息对象
func (c *CmdObjManager) GetProtoObjByCmdName(cmdName string) proto.Message {
	refType, ok := c.cmdNameProtoObjMap[cmdName]
	if !ok {
		logger.Error("unknown cmd name: %v", cmdName)
		return nil
	}
	protoObjInst := reflect.New(refType.Elem())
	protoObj := protoObjInst.Interface().(proto.Message)
	return protoObj
}

// GetCmdNameByProtoObj 通过消息对象获取操作名
func (c *CmdObjManager) GetCmdNameByProtoObj(protoObj proto.Message) string {
	cmdName, ok := c.protoObjCmdNameMap[reflect.TypeOf(protoObj)]
	if !ok {
		logger.Error("unknown cmd name: %v", cmdName)
		return ""
	}
	return cmdName
}
