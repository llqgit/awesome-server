package asinterface

type IPlugin interface {
	Start(server IServer)
	Stop()
}
