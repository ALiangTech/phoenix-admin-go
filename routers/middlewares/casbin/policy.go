package casbins

import "github.com/casbin/casbin/v2"

/*
* @desc 添加权限规则
* 这块是关联服务的接口的
* 如果接口地址修改 这边也要同步修改
 */

func AddPoliciesFxForApi(enforcer *casbin.Enforcer) error {
	policy := [][]string{
		{"user", "/user/*", "GET"}, // 用户管理模块
	}
	ok, err := enforcer.AddPolicies(policy)
	if !ok {
		return err
	}

	return enforcer.SavePolicy()
}
