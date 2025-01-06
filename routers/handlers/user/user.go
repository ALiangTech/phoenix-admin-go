package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"phoenix-go-admin/database"
	"phoenix-go-admin/routers/model/respond"
)

type User struct {
	Uid      string `json:"uid"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	RoleId   int    `json:"role_id"`
	Email    string `json:"email"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

// 根据当前登录用户 获取用户信息
func getUserInfo(ctx *gin.Context) {
	user, err := GetUserInfo(ctx)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, respond.Response{
		Data: user,
	})
}

func InitUserRouter(router *gin.RouterGroup) {
	router.GET("/user", getUserInfo)
}

// 根据uid 从数据库读取完整的用户信息
func (user *User) getUserInfoByUid() error {
	res := database.DB.Raw("SELECT uid,name,avatar,role_id,email,create_at,update_at FROM account WHERE uid = ?;", user.Uid).Scan(&user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// GetUserInfo 获取用户信息
func GetUserInfo(ctx *gin.Context) (User, error) {
	user := User{}
	uid, _ := ctx.Get("Uid") // 错误没有处理
	user.Uid = uid.(string)
	if err := user.getUserInfoByUid(); err != nil { // 错误没有处理
		return user, err
	}
	return user, nil
}
