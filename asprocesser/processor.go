package asprocesser

import "github.com/llqgit/awosome-server/asinterface"

// 创建一个通讯处理者（websocket）
func NewProcessor(t string) asinterface.IProcessor {
	switch t {
	case "websocket":
		return WebsocketProcessor{}
	}
	return nil
}
