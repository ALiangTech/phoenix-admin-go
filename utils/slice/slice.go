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

// Contains contains 判断一个字符串是否在一个字符串切片中。
func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// RemoveDuplicates 去掉二维数组中的重复项 根据二维数组的某一列的值进行去重
func RemoveDuplicates(slice [][]string, column int) [][]string {
	m := make(map[string]bool)
	var result [][]string
	for _, v := range slice {
		if _, ok := m[v[column]]; !ok {
			m[v[column]] = true
			result = append(result, v)
		}
	}
	return result
}
