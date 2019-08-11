package asnet

import (
	"fmt"
	"github.com/gorilla/websocket"
	_ "github.com/llqgit/awosome-server/asinterface"
	"sync"
)

var sid uint32 = 0 // 全局变量 session id

// 用户 session
type Session struct {
	Sid           uint32                     // session id （自增）
	Conn          *websocket.Conn            // websocket connection
	Property      map[string]interface{}     // session 属性存储
	PropertyLock  sync.RWMutex               // 读取自定义信息的锁（读写锁）
	Kick          bool                       // 是否被踢出
	MsgQueue      chan []byte                // 消息队列
	HeartBeatTime int64                      // 心跳时间
	Signal        chan bool                  // 关闭信号
	MsgHandler    func(s *Session, d []byte) // 消息处理者
}

// 创建一个 session （用户连接相关的信息）
func NewSession(conn *websocket.Conn, msgHandler func(s *Session, d []byte)) *Session {
	sid++
	session := &Session{
		Sid:        sid,
		Conn:       conn,
		Kick:       false,
		Property:   make(map[string]interface{}),
		MsgQueue:   make(chan []byte),
		MsgHandler: msgHandler,
	}
	return session
}

// 获取 session id
func (s *Session) GetSid() uint32 {
	return s.Sid
}

// 向 session 中设置自定义属性
func (s *Session) SetProperty(key string, value interface{}) {
	s.PropertyLock.Lock()
	defer s.PropertyLock.Unlock()
	// 如果值为 nil 则删除此 k，v
	if value == nil {
		delete(s.Property, key)
		return
	}
	s.Property[key] = value
}

// 读取 session 中的自定义属性
func (s *Session) GetProperty(key string) interface{} {
	s.PropertyLock.RLock()
	defer s.PropertyLock.RUnlock()
	if value, ok := s.Property[key]; ok {
		return value
	}
	return nil
}

// 发送消息到客户端（编码后的二进制数据）
func (s *Session) SendMessage(data []byte) error {
	if err := s.Conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
		return err
	}
	return nil
}

// 返回是否被踢出
func (s *Session) IsKicked() bool {
	return s.Kick
}

// 启动读操作
func (s *Session) StartReader() {
	//fmt.Println("[Reader Goroutine is running]")
	//defer fmt.Println(s.RemoteAddr().String(), "[conn Reader exit!]")
	//defer s.Stop()

	for {
		msgType, data, err := s.Conn.ReadMessage()
		if err != nil {
			s.Kick = true
			fmt.Printf("u^%v] receive error [%v]\n", s.Sid, err)
			return
		}
		if msgType == websocket.BinaryMessage {
			// 向上层传递接收的消息数据
			s.MsgHandler(s, data)
		}
	}

}

// 启动写操作
func (s *Session) StartWriter() {
	//defer s.Stop()
	// 从 session 的消息队列中读取消息，进行发送
	for {
		select {
		case data := <-s.MsgQueue:
			if err := s.SendMessage(data); err != nil {
				s.Kick = true
				fmt.Printf("u^%v] send message err [%v]\n", s.Sid, err)
				return
			}
		case stop := <-s.Signal:
			fmt.Println("session writer stop", stop)
			return
		}
	}
}

// 开始接受消息
func (s *Session) Start() {
	//1 开启用户从客户端读取数据流程的Goroutine
	go s.StartReader()
	//2 开启用于写回客户端数据流程的Goroutine
	go s.StartWriter()
	//按照用户传递进来的创建连接时需要处理的业务，执行钩子方法
	//s.TcpServer.CallOnConnStart(s)
}

// 结束 session
func (s *Session) Stop() {
	fmt.Println("session stop")
	s.Signal <- true
}

// 将需要发送的信息压入队列
func (s *Session) Send(data []byte) {
	if data == nil {
		fmt.Printf("session send msg should not be nil\n")
		return
	}
	// 如果没有被踢出，则正常添加消息
	if s.Kick == false {
		s.MsgQueue <- data
	}
}
