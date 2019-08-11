package group

import (
	"fmt"
	"github.com/llqgit/awosome-server/asnet"
	"github.com/llqgit/awosome-server/utils"
)

/**
一个 group 是 N 个用户的集合
推送消息时，会给 group 中的所有成员广播
*/

type Group struct {
	Id        string                    // 使用 uuid 字符串
	UserList  map[uint32]*asnet.Session // 玩家 session 列表(sid: session)
	UserIndex []uint32                  // 玩家 session 索引列表
	MaxSize   int                       // 最大容量
	MemberId  uint32                    // 成员 ID 的最大值
}

// 获取组成员数量
func (g *Group) GetSize() int {
	return len(g.UserList)
}

// 设置最大容量
func (g *Group) SetMaxSize(maxSize int) {
	g.MaxSize = maxSize
}

// 判断是否已经满员
func (g *Group) IsFull() bool {
	return g.GetSize() >= g.MaxSize
}

// 获取组成员
func (g *Group) GetMember(sid uint32) *asnet.Session {
	if member, ok := g.UserList[sid]; ok {
		return member
	}
	return nil
}

// 根据索引获取用户
func (g *Group) GetMemberByIndex(index int) *asnet.Session {
	if index >= g.Size() {
		return nil
	}
	fmt.Println(g.UserIndex)
	fmt.Println(g.UserList)
	sid := g.UserIndex[index]
	return g.GetMember(sid)
}

// 群发消息（包括组中的所有成员）
func (g *Group) SendMessage(data []byte) {
	for _, userSession := range g.UserList {
		userSession.Send(data)
	}
}

// 添加用户
func (g *Group) Add(session *asnet.Session) bool {
	// 超出最大成员个数限制，不能继续添加
	if g.IsFull() {
		return false
	}
	member := g.GetMember(session.Sid)
	if member != nil {
		return true
	}
	// 增加新的用户到组中
	g.MemberId++
	// 向 session 中设置 组成员 id
	session.SetProperty("MemberId", g.MemberId)    // 设置成员 id
	g.UserIndex = append(g.UserIndex, session.Sid) // 增加 sid 到索引列表
	g.UserList[session.Sid] = session
	return true
}

// 获取组的容量
func (g *Group) Size() int {
	return len(g.UserList)
}

// 移除用户
func (g *Group) Remove(sid uint32) {
	session := g.GetMember(sid)
	// 删除 session 中的 组成员 id
	session.SetProperty("MemberId", nil)
	utils.DeleteSlice(g.UserIndex, session.Sid) // 删除 sid 元素
	delete(g.UserList, sid)
}

// 移除所有用户
func (g *Group) RemoveAll() {
	for sid := range g.UserList {
		g.Remove(sid)
	}
}
