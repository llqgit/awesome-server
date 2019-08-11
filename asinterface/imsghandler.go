package asinterface

type IMsgHandler interface {
	SetApis(apis []IApi)           // 设置 api
	BeforeHandle(request IRequest) // 前置处理者
	Handle(request IRequest)       // 消息处理者
	AfterHandle(request IRequest)  // 后置处理者
}
