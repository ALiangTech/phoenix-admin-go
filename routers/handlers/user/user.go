package user

import "github.com/gin-gonic/gin"

/*
 * @desc 获取用户列表
 */
func pullUsers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "success",
	})
}

func InitUserRouter(router *gin.RouterGroup) {
	router.GET("/user", pullUsers)
}
