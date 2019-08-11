package utils

import (
	"github.com/satori/go.uuid"
)

// 获取一个 string 类型的 UUID
func GetUUID() string {
	id := uuid.Must(uuid.NewV4())
	return id.String()
}
