package mistakes

// 业务错误码

const (
	RoleExist    = 10001 //角色不存在
	ParamError   = 10002 // 参数存在问题
	SqlError     = 10003
	UserNotExist = 10004
	JwtError     = 10005
)

func StatusText(code int) string {
	switch code {
	case RoleExist:
		return "角色不存在"
	case ParamError:
		return "参数存在问题"
	case SqlError:
		return "sql执行错误"
	case UserNotExist:
		return "用户不存在"
	case JwtError:
		return "jwt生成失败"
	default:
		return ""
	}
}
