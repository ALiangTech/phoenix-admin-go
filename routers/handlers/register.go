package handlers

import (
	"phoenix-go-admin/routers/handlers/login"
	"phoenix-go-admin/routers/handlers/user"
	auth "phoenix-go-admin/routers/middlewares/authorization"
	casbins "phoenix-go-admin/routers/middlewares/casbin"

	"github.com/gin-gonic/gin"
)

/*
* @desc 接口分三种 需要登录 不需要登录 登录后的接口有分为有权限访问和无权限访问 （这边设计成登录后的接口全部需要权限访问，不在区分权限和无权限）
* 需要登录 登录后需要权限
* 需要登录 登录后不需要权限
* 不需要登录 直接访问
 */

func protectedApi(RouterGroup *gin.RouterGroup) {
	RouterGroup.Use(auth.Authorization(), casbins.CasbinCheck())
	// 需要登录才可以访问的接口
	user.InitUserRouter(RouterGroup)
	user.InitRoleRouter(RouterGroup)
}

func noProtectedApi(RouterGroup *gin.RouterGroup) {
	// 不需要登录即可访问的接口
	login.InitLoginRouter(RouterGroup)
}

func RegisterRouter(protectedApiGroup *gin.RouterGroup, noprotectedApiGroup *gin.RouterGroup) {
	protectedApi(protectedApiGroup)
	noProtectedApi(noprotectedApiGroup)
}
