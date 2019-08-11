package asinterface

type ISession interface {
	GetSid() uint32                            // 获取 session id
	SetProperty(key string, value interface{}) // 设置 session 自定义属性
	GetProperty(key string) interface{}        // 获取 session 自定义属性
	Start()                                    // 开始
	Stop()                                     // 停止
	IsKicked() bool                            // 是否被踢出
}
