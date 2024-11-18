package mistakes

// 业务错误码

const (
	RoleExist = 10001 //角色不存在
)

func StatusText(code int) string {
	switch code {
	case RoleExist:
		return "角色不存在"
	default:
		return ""
	}
}
