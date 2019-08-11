package asinterface

// 通讯的类型接口（tcp/udp/websocket）

type IProcessor interface {
	Listen(server IServer, ip string, port int, ipVersion string) error // 监听新的客户端接入
	Accept() (ISession, error)                                          // 接受新的客户端连接
	Resolve()                                                           // 处理客户端发来的数据
	Send()                                                              // 发送数据给客户端
}
