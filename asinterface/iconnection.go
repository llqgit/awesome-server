package asinterface

type IConnection interface {
	SendMsg(request IRequest)
}
