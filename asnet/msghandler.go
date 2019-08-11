package asnet

import "github.com/llqgit/awosome-server/asinterface"

// 消息处理者，主要负责接受客户端传来的消息，进行处理
type MsgHandler struct {
	Apis           map[uint32]asinterface.IApi // 存放每个MsgId 所对应的处理方法的map属性
	WorkerPoolSize uint32                      // 业务工作Worker池的数量
	MsgQueue       []chan asinterface.IRequest
}

// 创建一个消息处理者
func NewMsgHandler(workerPoolSize uint32) *MsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]asinterface.IApi),
		WorkerPoolSize: workerPoolSize,
		MsgQueue:       make([]chan asinterface.IRequest, workerPoolSize),
	}
}

// 设置 api
func (h MsgHandler) SetApis(apis []asinterface.IApi) {
	h.Apis = make(map[uint32]asinterface.IApi)
	for _, api := range apis {
		h.Apis[api.GetApiId()] = api
	}
}

// 处理
func (h MsgHandler) DoHandle(request asinterface.IRequest) {
	h.BeforeHandle(request)
	h.Handle(request)
	h.AfterHandle(request)
}

// 前置消息处理者
func (h MsgHandler) BeforeHandle(request asinterface.IRequest) {}

// 消息处理者
func (h MsgHandler) Handle(request asinterface.IRequest) {
	apiID := request.GetApi()
	api, ok := h.Apis[apiID]
	if !ok {
		return
	}
	api.Handler(request.GetSession(), api.GetPayload())
}

// 后置消息处理者
func (h MsgHandler) AfterHandle(request asinterface.IRequest) {}
