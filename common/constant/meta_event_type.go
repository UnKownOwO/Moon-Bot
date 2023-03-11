package constant

const (
	MetaEventTypeLifecycle               = "lifecycle" // 生命周期
	MetaEventTypeLifecycleSubTypeConnect = "connect"   // 客户端连接
	MetaEventTypeLifecycleSubTypeEnable  = "enable"    // 客户端启用
	MetaEventTypeLifecycleSubTypeDisable = "disable"   // 客户端禁用

	MetaEventTypeHeartbeat = "heartbeat" // 心跳包
)
