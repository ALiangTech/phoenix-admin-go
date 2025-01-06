package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"phoenix-go-admin/database"
	casbins "phoenix-go-admin/routers/middlewares/casbin"
	"phoenix-go-admin/routers/model/respond"
	"phoenix-go-admin/utils/slice"
)

type Role struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CasbinRole string `json:"casbin_role"`
	CreateAt   string `json:"create_at"`
	UpdateAt   string `json:"update_at"`
}

// RoleTree 树节点结构
type RoleTree struct {
	Label    string     `json:"label"`
	Value    string     `json:"value"`
	Children []RoleTree `json:"children"`
}

var roleTree = []RoleTree{
	{
		Label: "用户管理",
		Value: "user",
		Children: []RoleTree{
			{
				Label: "账号管理",
				Value: "user_account",
				Children: []RoleTree{
					{
						Label:    "账号添加",
						Value:    "user_account_add",
						Children: []RoleTree{},
					},
					{
						Label:    "账号编辑",
						Value:    "user_account_edit",
						Children: []RoleTree{},
					},
					{
						Label:    "账号删除",
						Value:    "user_account_delete",
						Children: []RoleTree{},
					},
				},
			},
			{
				Label: "角色管理",
				Value: "user_role",
				Children: []RoleTree{
					{
						Label:    "角色添加",
						Value:    "user_role_add",
						Children: []RoleTree{},
					},
					{
						Label:    "角色编辑",
						Value:    "user_role_edit",
						Children: []RoleTree{},
					},
					{
						Label:    "角色删除",
						Value:    "user_role_delete",
						Children: []RoleTree{},
					},
				},
			},
		},
	},
}

// filterTree 根据 codes 过滤树结构
func filterTree(tree []RoleTree, codes map[string]bool) []RoleTree {
	var result []RoleTree
	for _, node := range tree {
		// 过滤子节点
		filteredChildren := filterTree(node.Children, codes)

		// 检查当前节点是否在 codes 中或有子节点保留下来
		if codes[node.Value] || len(filteredChildren) > 0 {
			result = append(result, RoleTree{
				Label:    node.Label,
				Value:    node.Value,
				Children: filteredChildren,
			})
		}
	}
	return result
}

// 生成角色树
// 流程查询当前用户拥有权限 => 过滤掉没有权限的节点 => 返回节点
func getRoleTree(ctx *gin.Context) {
	// 查询当前用户拥有角色
	user, userError := GetUserInfo(ctx)
	if userError != nil {
		return
	}
	role, roleError := getRoleInfo(user.RoleId)
	if roleError != nil {
		return
	}
	filteredPolicy, casbinError := casbins.E.GetRolesForUser(role.CasbinRole)
	if casbinError != nil {
		return
	}
	codes := slice.ToMap(filteredPolicy)
	fmt.Println(filterTree(roleTree, codes))
	ctx.JSON(http.StatusOK, respond.Response{
		Data: filterTree(roleTree, codes),
	})
}

func InitRoleRouter(router *gin.RouterGroup) {
	router.GET("/user/role/tree", getRoleTree)
}

// getRoleInfo 根据角色ID 查询角色详情
func getRoleInfo(roleId int) (Role, error) {
	var role Role
	res := database.DB.Raw("select * from role where id = ?", roleId).Scan(&role)
	if res.RowsAffected == 0 {
		return role, fmt.Errorf("角色类型不存在")
	}
	if res.Error != nil {
		return role, res.Error
	}
	return role, nil
}
