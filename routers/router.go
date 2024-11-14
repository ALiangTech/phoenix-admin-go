package routers

import (
	"phoenix-go-admin/config/env"
	"phoenix-go-admin/routers/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	protectedApiGroup := router.Group("api/v1/protected")
	noProtectedApiGroup := router.Group("api/v1")
	handlers.RegisterRouter(protectedApiGroup, noProtectedApiGroup)
	router.Run(env.Config.HTTP_PORT)
}
