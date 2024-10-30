package login

import "github.com/gin-gonic/gin"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 处理登录
func login(ctx *gin.Context) {
	// 从请求体中获取参数
	loginRequest := LoginRequest{}
	if err := ctx.ShouldBind(&loginRequest); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}
	ctx.JSON(400, gin.H{"error": "Invalid request data"})
}

func InitLoginRouter(router *gin.RouterGroup) {
	router.POST("login", login)
}
