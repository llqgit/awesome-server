package asinterface

type IApi interface {
	GetApiId() uint32                         // 获取 api id
	GetPayload() []byte                       // 获取二进制数据
	Handler(session ISession, payload []byte) // 获取处理者
}
