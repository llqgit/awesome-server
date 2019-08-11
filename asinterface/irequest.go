package asinterface

type IRequest interface {
	GetSession() ISession
	GetData() []byte
	GetApi() uint32
}
