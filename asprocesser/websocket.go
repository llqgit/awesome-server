package asprocesser

import (
	"fmt"
	"github.com/llqgit/awosome-server/asinterface"
	"github.com/llqgit/awosome-server/asnet"
	"golang.org/x/net/websocket"
	"net/http"
)

// websocket 的通讯协议
type WebsocketProcessor struct {
	Server *asinterface.IServer
}

// 监听新的客户端接入
func (p WebsocketProcessor) Listen(server asinterface.IServer, ip string, port int, ipVersion string) error {
	p.Server = &server
	http.Handle("/", websocket.Handler(p.Handler))
	if err := http.ListenAndServe(fmt.Sprintf("%v:%v", ip, port), nil); err != nil {
		return err
	}
	return nil
}

// 处理者
func (p WebsocketProcessor) Handler(ws *websocket.Conn) {
	//3.1 阻塞等待客户端建立新的连接请求
	//fmt.Println("Get conn remote addr = ", conn.RemoteAddr().String())

	//3.2 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
	//if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
	//	conn.Close()
	//	continue
	//}

	//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
	dealSession := asnet.NewSession(ws, nil)

	//3.4 启动当前链接的处理业务
	go dealSession.Start()
}

// 接受新的客户端连接
func (p WebsocketProcessor) Accept() (asinterface.ISession, error) {
	return nil, nil
}

// 处理客户端发来的数据
func (p WebsocketProcessor) Resolve() {

}

// 发送数据给客户端
func (p WebsocketProcessor) SendMessage(data []byte) {

}
