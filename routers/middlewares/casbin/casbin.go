package casbins

import (
	"database/sql"
	"fmt"
	"net/http"
	"phoenix-go-admin/config/env"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	casbinpgadapter "github.com/cychiuae/casbin-pg-adapter"
	"github.com/gin-gonic/gin"
)

var casbinModel = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m =  g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act) || r.sub == "root"
`
var E *casbin.Enforcer

func init() {
	// connectionString := "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable"
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", env.Config.DB_USER, env.Config.DB_PASSWORD, env.Config.DB_IP, env.Config.DB_NAME)
	fmt.Println("connectionString", connectionString)
	db, err := sql.Open("postgres", connectionString)
	fmt.Println("err", err)
	if err != nil {
		panic(err)
	}
	tableName := "casbin"
	adapter, err := casbinpgadapter.NewAdapter(db, tableName)
	if err != nil {
		panic(err)
	}
	m, _ := model.NewModelFromString(casbinModel)
	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		panic(err)
	}
	E = enforcer

}

/**
* @desc 检查用户是否拥有权限访问该接口
**/
func CasbinCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		sub, exists := c.Get("sub")
		if exists {
			// sub := "alice" // the user that wants to access a resource.
			obj := c.Request.URL.Path // the resource that is going to be accessed.
			act := c.Request.Method   // the operation that the user performs on the resource.

			ok, err := E.Enforce(sub, obj, act)
			fmt.Println(ok, err, obj, act, "sss")
			if ok {
				c.Next()
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 500,
					"msg":  "暂无权限",
				})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "角色不存在",
			})
			c.Abort()
			return
		}
	}
}
