package role

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"phoenix-go-admin/routers/handlers/user/role/common"
	"phoenix-go-admin/routers/handlers/user/role/tree"
	casbins "phoenix-go-admin/routers/middlewares/casbin"
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
	router.GET("/user/role/add", addRole)
}

// 创建角色
// 获取参数 => 插入数据库 => 返回结果
// 针对codes 还得做一个校验确保codes 属于当前用户拥有的权限码: 防止通过接口胡乱传递codes导致拥有了其他权限

type addParams struct {
	Name  string   `json:"name" binding:"required,min=1,max=30"`
	Codes []string `json:"codes" binding:"required,unique,dive,min=1,max=80"`
}
type addParamsValidate struct {
	addParams
	isValidate bool
}

// 插入数据库

func (receiver *addParamsValidate) validateCodes(ctx *gin.Context) error {
	roleDetail, err := common.GetRoleDetailByContext(ctx)
	if err != nil {
		return err
	}
	isValidate, err := common.ValidateCodes(receiver.Codes, roleDetail.CasbinRole)
	if err != nil {
		return err
	}
	receiver.isValidate = isValidate
	return nil
}

func addRole(ctx *gin.Context) {
	var role addParamsValidate
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusOK, respond.Response{
			Code: 400,
			Msg:  err.Error(),
		})
		return
	}
	if err := role.validateCodes(ctx); err != nil {
		ctx.JSON(http.StatusOK, respond.Response{
			Code: 400,
			Msg:  err.Error(),
		})
		return
	}
	err := casbins.BuildRole(role.Codes, casbins.E)
	if err != nil {
		ctx.JSON(http.StatusOK, respond.Response{
			Code: 500,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, respond.Response{
		Data: role,
	})
}
