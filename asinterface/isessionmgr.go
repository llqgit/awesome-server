package asinterface

type ISessionMgr interface {
	GetSession(sid uint32) (ISession, bool) // 根据 sid 获取一个 session
	AddSession(session ISession)            // 添加一个新的 session 到管理器
	RemoveSession(sid uint32)               // 清除一个 session
	GetSessionCount(session ISession) int   // 获取当前 session 数量
	KickOne(sid uint32)                     // 踢出一个玩家 session
	KickAll()                               // 踢出所有玩家 session
}
