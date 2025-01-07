package casbins

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"phoenix-go-admin/utils/slice"
)

/*
* @desc 添加权限规则
* 这块是关联服务的接口的
* 如果接口地址修改 这边也要同步修改
 */

func AddPoliciesFxForApi(enforcer *casbin.Enforcer) error {
	// 专用权限 不仅需要登录还需要开启对应权限才可以访问
	policy := [][]string{
		{"user", "_", "GET"},                                // 用户管理 "_"的话一般是抽象层 没有对应接口, 方便前端使用
		{"user_account", "/user/user/*", "GET"},             // 用户管理_账号管理
		{"user_account_add", "/user/user/add/*", "POST"},    // 用户管理_账号管理_添加账号
		{"user_account_edit", "/user/user/edit/*", "PATCH"}, // 用户管理_账号管理_编辑账号
		{"user_account_delete", "/user/user/*", "DELETE"},   // 用户管理_账号管理_删除账号
		{"user_role", "/user/role/*", "GET"},                // 用户管理_角色管理
		{"user_role_add", "/user/role/add", "POST"},         // 用户管理_角色管理_添加角色
		{"user_role_edit", "/user/role/edit", "PATCH"},      // 用户管理_角色管理_修改角色
		{"user_role_delete", "/user/role/delete", "DELETE"}, // 用户管理_角色管理_删除角色
	}
	// 通用权限 只需要登录即可访问(属于用户登录了就可以访问)
	common := [][]string{
		{"common", "/user", "GET"}, // 获取当前登录用户信息
	}
	policies := append(policy, common...)
	ok, err := enforcer.AddPoliciesEx(policies) // 不存在的规则会被添加
	// 手动创建一个admin 角色
	buildAdminRole(policies, enforcer)
	if !ok {
		return err
	}

	return enforcer.SavePolicy()
}

// 基于policies 创建一个admin角色
func buildAdminRole(policies [][]string, enforcer *casbin.Enforcer) {
	var groupingPolicies [][]string
	for _, sub := range policies {
		groupingPolicies = append(groupingPolicies, []string{"admin", sub[0]})
	}
	ok, err := enforcer.AddGroupingPoliciesEx(groupingPolicies)
	if !ok {
		panic(err)
	}
}

// BuildRole 根据权限codes 创建一个角色
// 角色名称生成规则role_1
func BuildRole(codes []string, enforcer *casbin.Enforcer) error {
	namedGroupingPolicy, _ := enforcer.GetNamedGroupingPolicy("g")       // 获取所有角色 [[name,1],[name,2]]
	namedGroupingPolicy = slice.RemoveDuplicates(namedGroupingPolicy, 0) // 去重 [[name,1],[name,2]] => [[name,1]]
	roleName := fmt.Sprintf("role_%d", len(namedGroupingPolicy)+1)
	var groupingPolicies [][]string
	for _, sub := range codes {
		groupingPolicies = append(groupingPolicies, []string{roleName, sub})
	}
	_, err := enforcer.AddGroupingPolicies(groupingPolicies)
	return err
}
