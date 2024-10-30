package routers

import (
	"phoenix-go-admin/routers/handlers/login"

	"github.com/gin-gonic/gin"
)

var ApiRouterGroup *gin.RouterGroup

func InitRouter() {
	router := gin.Default()
	ApiRouterGroup = router.Group("api")
	login.InitLoginRouter(ApiRouterGroup)
	router.Run(":9000")
}
