package asplugin

import "github.com/llqgit/awosome-server/asinterface"

type StandardProtocol struct {
}

// 运行
func (p StandardProtocol) Do(server asinterface.IServer) {
	apis := make([]asinterface.IApi, 0)
	msgHandler := server.GetMsgHandler()
	msgHandler.SetApis(apis)
}
