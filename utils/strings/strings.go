package utils

// 从字符串中获取排除 prefix 前缀的 path
func GetStringWithoutPrefix(path string, prefix string) string {
	return path[len(prefix):]
}
