package login

import (
	"net/http"
	"phoenix-go-admin/routers/model/respond"

	"github.com/gin-gonic/gin"
)

type request struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 处理登录
func login(ctx *gin.Context) {
	// 从请求体中获取参数
	loginRequest := request{}
	err := ctx.ShouldBind(&loginRequest)
	if err != nil {
		var msg []interface{}
		msg = append(msg, "参数存在问题")
		ctx.JSON(http.StatusBadRequest, respond.Response{
			Data: nil,
			Err:  err.Error(),
			Msg:  msg,
		})
		return
	}
	ctx.JSON(400, gin.H{"error": "Invalid request data"})
}

func InitLoginRouter(router *gin.RouterGroup) {
	router.POST("login", login)
}
