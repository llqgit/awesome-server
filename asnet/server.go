package asnet

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// 服务器
type Server struct {
	Name        string                 // 服务器名称
	IP          string                 // 服务器 IP
	IPVersion   string                 // IP 类型 v4 或 其他
	Port        int                    // 服务绑定的端口号
	SessionList []*Session             // session 管理器（现在就是一个数组）
	OnConnect   func(*Session)         // 连接成功回调
	OnClosed    func(*Session)         // 连接关闭回调
	MsgHandler  func(*Session, []byte) // 消息回调
}

// 创建一个新的服务器对象
func NewServer(name string, ip string, port int, msgHandler func(*Session, []byte)) *Server {
	return &Server{
		Name:        name,
		IPVersion:   "tcp4",
		IP:          ip,
		Port:        port,
		SessionList: []*Session{},
		MsgHandler:  msgHandler,
	}
}

// 开始服务器
func (s *Server) Start() {
	// 开启一个 go协程 去做服务端 listener 业务
	go func() {
		// 默认路由为根目录
		http.HandleFunc("/", s.Handle)
		// 启动服务器监听（会阻塞到这里）
		if err := http.ListenAndServe(fmt.Sprintf("%v:%v", s.IP, s.Port), nil); err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}
	}()
	// 打印正在开始监听
	fmt.Println("start Awesome Server  ", s.Name, " success, now listening...")
}

// 运行服务器
func (s *Server) Serve() {
	s.Start()
	select {} // 保证主线程不结束
}

// 处理接收到的信息回调
func (s *Server) NetMsgHandler(session *Session, data []byte) {
	// TODO 加密解密数据
	//netMsg := &protocol.NetMsg{}
	//err := proto.Unmarshal(data, netMsg)
	//if err != nil {
	//	fmt.Println("err", err)
	//	return
	//}
	//fmt.Printf("s.MsgHandler != nil %+v\n", s.MsgHandler)
	// 上传到上层解析
	if s.MsgHandler != nil {
		s.MsgHandler(session, data)
	}
}

// 将 http 协议升级为 websocket 协议的参数（gorilla/websocket 要这么用）
var upgradeOpt = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 处理新进入的链接
func (s *Server) Handle(w http.ResponseWriter, r *http.Request) {
	conn, err := upgradeOpt.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	// 创建一个 session 会话，包含一个 conn 链接参数，以及一个处理读取信息的回调
	session := NewSession(conn, s.NetMsgHandler)

	// 加入 session 列表
	s.SessionList = append(s.SessionList, session)
	fmt.Printf("session length [%+v] current connect sid [%v]\n", len(s.SessionList), session.Sid)

	// 连接成功回调
	if s.OnConnect != nil {
		s.OnConnect(session)
	}

	// 执行 session 的 start 方法，开启 read 和 write go协程
	session.Start()
}

// 设置连接成功回调
func (s *Server) SetOnConnect(cb func(*Session)) {
	s.OnConnect = cb
}

// 设置连接关闭回调
func (s *Server) SetOnClosed(cb func(*Session)) {
	s.OnClosed = cb
}

// 设置信息处理器
func (s *Server) SetMsgHandler(cb func(*Session, []byte)) {
	s.MsgHandler = cb
}
