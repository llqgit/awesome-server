package heartbeat

import (
	"github.com/llqgit/awosome-server/asnet"
	"time"
)

// 心跳功能

type HeartBeat struct {
	UserQueue []*asnet.Session
}

func (h *HeartBeat) Check() {
	now := time.Now().UnixNano()
	for _, session := range h.UserQueue {
		// 如果 session 的心跳时间 > 5 秒
		if now-session.HeartBeatTime > 5*1e9 {
			session.Stop()
		}
	}
}
