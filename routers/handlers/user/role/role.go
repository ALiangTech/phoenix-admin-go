package role

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"phoenix-go-admin/routers/handlers/user/role/tree"
	"phoenix-go-admin/routers/model/respond"
)

// getRoleTree 获取角色树 API
func getRoleTree(ctx *gin.Context) {
	roleTree, err := tree.BuildUserTree(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, respond.Response{
			Code: 500,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, respond.Response{
		Data: roleTree,
	})
}

func InitRoleRouter(router *gin.RouterGroup) {
	router.GET("/user/role/tree", getRoleTree)
}
