package routers

import (
	"phoenix-go-admin/config/env"
	"phoenix-go-admin/routers/handlers"
	utils "phoenix-go-admin/utils/strings"

	"github.com/gin-gonic/gin"
)

// 接口前缀
var baseApi = "/api/v1"
var Protected = "/api/v1/protected"

func InitRouter() {
	router := gin.Default()
	protectedApiGroup := router.Group(Protected)
	protectedApiGroup.Use(createCasbinObj())
	noProtectedApiGroup := router.Group(baseApi)
	handlers.RegisterRouter(protectedApiGroup, noProtectedApiGroup)
	err := router.RunTLS(env.Config.HTTP_PORT, "./config/ca/PHONENIX_Root_CA.crt", "./config/ca/phoenix.pem")
	if err != nil {
		return
	}
}

// 通过gin中间件的方式统一设置请求的obj
func createCasbinObj() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("obj", utils.GetStringWithoutPrefix(c.Request.URL.Path, Protected))
	}
}
