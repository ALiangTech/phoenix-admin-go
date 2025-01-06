package slice

// ToMap 将字符串切片转换为一个映射，其中切片中的每个元素作为映射中的一个键。
// 这个函数的主要目的是为了快速去重和查找。
func ToMap(slice []string) map[string]bool {
	m := make(map[string]bool, len(slice))
	for _, v := range slice {
		m[v] = true
	}
	return m
}
