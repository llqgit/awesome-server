package utils

// 删除 slice 的一个元素
func DeleteSlice(a []uint32, num uint32) []uint32 {
	j := uint32(0)
	for _, val := range a {
		if val == num {
			a[j] = val
			j++
		}
	}
	return a[:j]
}
