package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"phoenix-go-admin/config/env"
	"phoenix-go-admin/database"
	"phoenix-go-admin/routers/model/respond"
	jwtToken "phoenix-go-admin/utils/jwt_token"
	"phoenix-go-admin/utils/mistakes"
)

type credential struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type user struct {
	Uid    string
	RoleId int
}
type token struct {
	Jwt string `json:"jwt" binding:"required"`
}

type loginContext struct {
	Credential credential
	User       user
	Token      token
	Ctx        *gin.Context
}

// 处理登录
// 正常流程:
// 获取登录凭证
// 根据登录凭证查询用户唯一ID
// 根据唯一ID生成token
// 返回token

func login(ctx *gin.Context) {
	loginContext := loginContext{Ctx: ctx}
	if err := loginContext.getCredential(); err != nil {
		mistakes.HandleErrorResponse(ctx, mistakes.ParamError, err)
		return
	}
	if err := loginContext.getUidByCredential(); err != nil {
		mistakes.HandleErrorResponse(ctx, mistakes.UserNotExist, err)
		return
	}
	if err := loginContext.generateToken(); err != nil {
		mistakes.HandleErrorResponse(ctx, mistakes.JwtError, err)
		return
	}

	ctx.JSON(http.StatusOK, respond.Response{
		Data: map[string]interface{}{
			"jwt": loginContext.Token.Jwt,
		},
	})
}

func InitLoginRouter(router *gin.RouterGroup) {
	router.POST("login", login)
}

// 从请求中获取登录凭证
func (context *loginContext) getCredential() error {
	err := context.Ctx.ShouldBind(&context.Credential)
	return err
}

// 从数据库根据登录凭证查询是否存在这个用户 存在就返回用户uid/role_id
func (context *loginContext) getUidByCredential() error {
	// 白嫖的pgsql 不支持crypt 先这样
	//res := database.DB.Raw("SELECT uid FROM user WHERE name = ? AND password = crypt(?, password);", user.Username, user.Password)
	res := database.DB.Raw("SELECT uid, role_id FROM account WHERE name = ?;", context.Credential.Username).Scan(&context.User)
	if res.Error != nil {
		return res.Error
	}
	if context.User.Uid == "" {
		return fmt.Errorf("用户不存在")
	}
	return res.Error
}

// 根据用户唯一ID 生成Token
func (context *loginContext) generateToken() error {
	jwt, err := jwtToken.GenerateJwt(env.Config.JWT_SECRET, context.User.Uid)
	if err != nil {
		return err
	}
	if jwt == "" {
		return fmt.Errorf("生成token失败")
	}
	context.Token.Jwt = jwt
	return err
}
