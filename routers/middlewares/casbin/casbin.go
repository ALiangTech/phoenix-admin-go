package casbins

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"net/http"
	"phoenix-go-admin/config/env"
	"phoenix-go-admin/utils/mistakes"

	pgadapter "github.com/casbin/casbin-pg-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/gin-gonic/gin"
)

// p 是角色
// g 是用户与角色的关系
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
m =  g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)
`
var E *casbin.Enforcer

func init() {
	// connectionString := "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable"
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", env.Config.DB_USER, env.Config.DB_PASSWORD, env.Config.DB_IP, env.Config.DB_NAME)
	fmt.Println("connectionString", connectionString)
	opts, _ := pg.ParseURL(connectionString)
	db := pg.Connect(opts)
	defer db.Close()
	adapter, err := pgadapter.NewAdapterByDB(db, pgadapter.WithTableName("casbin"))
	fmt.Println("err", err)
	if err != nil {
		panic(err)
	}
	m, _ := model.NewModelFromString(casbinModel)
	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		panic(mistakes.NewError("构建enforcer失败", err))
	}
	err = AddPoliciesFxForApi(enforcer)
	if err != nil {
		panic(mistakes.NewError("添加策略失败", err))
	}
	E = enforcer
	err = E.LoadPolicy()
	if err != nil {
		panic(mistakes.NewError("策略加载失败", err))
	}
	allSubjects, _ := E.GetAllSubjects()
	fmt.Println("allSubjects")
	fmt.Println(allSubjects)
}

func CasbinCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		sub, roleExists := c.Get("Sub")
		obj, resourceExists := c.Get("Obj")
		act := c.Request.Method
		if roleExists && resourceExists {
			allowed, err := E.Enforce(sub, obj, act)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": 500,
					"msg":  "权限检查失败" + err.Error()})
				c.Abort()
				return
			}
			if allowed {
				c.Next()
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 403,
					"msg":  "没有权限访问该接口",
				})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": mistakes.RoleExist,
				"msg":  mistakes.StatusText(mistakes.RoleExist),
			})
			c.Abort()
			return
		}
	}
}
