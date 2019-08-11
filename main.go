package main

import (
	"fmt"
	"github.com/llqgit/awosome-server/asnet"
	protocol "knifeServer/protocol/go"
	"math/rand"
)

const REQ_LOGIN = 1       // 请求登录
const ACK_LOGIN = 2       // 回复登录
const REQ_LOGOUT = 3      // 请求注销
const REQ_JOIN_ROOM = 4   // 请求加入房间
const ACK_JOIN_ROOM = 5   // 回复加入房间
const NTF_ROOM_INFO = 6   // 通知房间信息
const NTF_GAME_START = 7  // 通知开始游戏
const REQ_HIT = 8         // 请求射击
const NTF_GAME_OVER = 9   // 通知游戏结束
const NTF_GAME_FRAME = 10 // 通知游戏帧同步

// 连接成功
func OnConnect(session *asnet.Session) {
	fmt.Println(".... on connect ....", session.Sid)
	randomName := fmt.Sprintf("hello_%v", rand.Intn(100))
	session.SetProperty("uid", randomName)
}

// 连接关闭
func OnClosed(session *asnet.Session) {

}

// 服务器收到的消息处理者
func MsgHandler(session *asnet.Session, msg *protocol.NetMsg) {
	name := session.GetProperty("uid")
	fmt.Printf("name-[%v]\n", name)
	fmt.Printf("recevie msg: [%+v]\n", msg)
	//switch msg.Api {
	//case REQ_JOIN_ROOM:
	//	var payload protocol.ReqJoinRoom
	//	err := proto.Unmarshal(msg.Payload, &payload)
	//	if err != nil {
	//		return
	//	}
	//	controller.ReqJoinRoom(session, payload)
	//}
}

func main() {
	fmt.Println("hello world")
	s := asnet.NewServer("[server name]", "127.0.0.1", 5600, MsgHandler)

	s.SetOnConnect(OnConnect)
	s.SetOnClosed(OnClosed)
	//s.SetMsgHandler(MsgHandler)
	//s.Use(asplugin.StandardProtocol{}) // 使用标准协议模版
	//s.Use(asplugin.WorkerPool{})       // 使用工作池机制

	//s.Use(config)          // 使用配置文件

	s.Serve()
}
