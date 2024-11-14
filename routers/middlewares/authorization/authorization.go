// gin 中间件 校验请求头是否携带Authorization
package auth

import (
	"fmt"
	"net/http"
	"phoenix-go-admin/config/env"
	jwtToken "phoenix-go-admin/utils/jwt_token"

	"github.com/gin-gonic/gin"
)

var secret string = env.Config.JWT_SECRET

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "请求头中Authorization为空",
			})
			c.Abort()
			return
		}
		// 获取到了Authorization
		// 解析 获取用户唯一id
		clamis, err := jwtToken.GetJwtsClamis(secret, auth)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		fmt.Println(clamis) // 获取了唯一id
		// 根据唯一id 获取用户角色
		c.Set("sub", "hh")
	}
}
