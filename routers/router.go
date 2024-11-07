package routers

import (
	"phoenix-go-admin/config/env"
	"phoenix-go-admin/routers/handlers/login"

	"github.com/gin-gonic/gin"
)

var ApiRouterGroup *gin.RouterGroup

func InitRouter() {
	router := gin.Default()
	ApiRouterGroup = router.Group("api")
	login.InitLoginRouter(ApiRouterGroup)
	router.Run(env.Config.HTTP_PORT)
	println(enforcer.Enforce("admin", "api/login", "POST"))
}
