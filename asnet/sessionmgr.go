package asnet

import (
	"github.com/llqgit/awosome-server/asinterface"
	"sync"
)

// session 管理器
type SessionMgr struct {
	SessionMap map[uint32]asinterface.ISession
	Lock       sync.RWMutex
}

// 创建一个 session 管理器
func NewSessionMgr() SessionMgr {
	m := SessionMgr{}
	go m.GC()
	return m
}

func (s SessionMgr) GC() {
	for sid, session := range s.SessionMap {
		if session.IsKicked() {
			delete(s.SessionMap, sid)
		}
	}
}

// 通过 sid 获取 session
func (s SessionMgr) GetSession(sid uint32) (asinterface.ISession, bool) {
	if session, ok := s.SessionMap[sid]; ok {
		return session, true
	} else {
		return nil, false
	}
}

// 添加 session
func (s SessionMgr) AddSession(session asinterface.ISession) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.SessionMap[session.GetSid()] = session
}

// 删除 session
func (s SessionMgr) RemoveSession(sid uint32) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	if session, ok := s.GetSession(sid); ok {
		session.Stop()
	}
	delete(s.SessionMap, sid)
}

// 获取当前 session 数量
func (s SessionMgr) GetSessionCount(session asinterface.ISession) int {
	return len(s.SessionMap)
}

// 踢出一个玩家 session
func (s SessionMgr) KickOne(sid uint32) {
	s.RemoveSession(sid)
}

// 踢出所有玩家 session
func (s SessionMgr) KickAll() {
	for sid := range s.SessionMap {
		s.KickOne(sid)
	}
}
