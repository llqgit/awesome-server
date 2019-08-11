package group

import (
	"github.com/llqgit/awosome-server/asnet"
	"github.com/llqgit/awosome-server/utils"
)

/**
组管理器（房间管理器）
负责管理同一类型房间的集合
*/

// 组管理器（全局唯一，暂时不需要ID）
type Manager struct {
	GroupList     map[string]*Group // 组列表
	MaxMemberSize int               // 每个组的最大成员数量
}

// 创建一个新的组管理器（每个房间最大成员数量）
func NewGroupManager(memberSize int) *Manager {
	return &Manager{
		GroupList:     make(map[string]*Group),
		MaxMemberSize: memberSize, // 默认每组人数
	}
}

// 设置最大的每组人数（每个房间的最大人数）
func (m *Manager) SetMaxMemberSize(size int) {
	m.MaxMemberSize = size
}

// 创建一个组（最大成员数量继承组管理器的设置）
func (m *Manager) NewGroup() *Group {
	group := &Group{
		Id:        utils.GetUUID(),                    // 组 ID
		UserList:  make(map[uint32]*asnet.Session, 0), // 组的 session list
		UserIndex: make([]uint32, 0),                  // 组的 sid 索引
		MaxSize:   m.MaxMemberSize,                    // 组的最大成员数量
		MemberId:  0,                                  // 成员ID 从 0 开始递增
	}
	return group
}

// 获取一个组
func (m *Manager) GetGroup(groupId string) *Group {
	if group, ok := m.GroupList[groupId]; ok {
		return group
	}
	return nil
}

// 获取一个未满组（有坐的房间）
func (m *Manager) GetGroupHaveSeat() *Group {
	for _, group := range m.GroupList {
		if group.IsFull() == false {
			return group
		}
	}
	newGroup := m.NewGroup()
	m.AddGroup(newGroup)
	return newGroup
}

// 添加一个组（返回是否新添加成功）
func (m *Manager) AddGroup(group *Group) bool {
	if _, ok := m.GroupList[group.Id]; ok {
		return false
	}
	m.GroupList[group.Id] = group
	return true
}

// 删除一个组
func (m *Manager) DeleteGroup(groupId string) {
	g := m.GetGroup(groupId)
	if g != nil {
		g.RemoveAll() // 调用 组 的清除方法
	}
	delete(m.GroupList, groupId)
}

// 获取已有的组的数量
func (m *Manager) GetGroupCount() int {
	return len(m.GroupList)
}
