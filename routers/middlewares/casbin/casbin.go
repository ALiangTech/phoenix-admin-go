package enforcer

import (
	"database/sql"
	"github.com/casbin/casbin/v2"
	"github.com/cychiuae/casbin-pg-adapter"
    "github.com/casbin/casbin/v2/model"

)

var casbinModel = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act) || r.sub == "root"
`
var Enforcer *casbin.Enforcer

func init() {
	connectionString := "postgresql://postgres:@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
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
	Enforcer = enforcer

}
