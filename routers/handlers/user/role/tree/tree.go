package tree

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"phoenix-go-admin/database"
	"phoenix-go-admin/routers/handlers/user/account"
	"phoenix-go-admin/routers/handlers/user/role/types"
	casbins "phoenix-go-admin/routers/middlewares/casbin"
	"phoenix-go-admin/utils/slice"
)

type Node struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Children []Node `json:"children"`
}

// 权限树
var tree = []Node{
	{
		Label: "用户管理",
		Value: "user",
		Children: []Node{
			{
				Label: "账号管理",
				Value: "user_account",
				Children: []Node{
					{
						Label: "账号添加",
						Value: "user_account_add",
					},
					{
						Label: "账号编辑",
						Value: "user_account_edit",
					},
					{
						Label: "账号删除",
						Value: "user_account_delete",
					},
				},
			},
			{
				Label: "角色管理",
				Value: "user_role",
				Children: []Node{
					{
						Label: "角色添加",
						Value: "user_role_add",
					},
					{
						Label: "角色编辑",
						Value: "user_role_edit",
					},
					{
						Label: "角色删除",
						Value: "user_role_delete",
					},
				},
			},
		},
	},
}

// filterTree 根据权限码过滤树结构
func filterTree(tree []Node, codes map[string]bool) []Node {
	result := []Node{} // 创建一个空的结果切片
	for _, node := range tree {
		// 过滤子节点
		filteredChildren := filterTree(node.Children, codes)

		// 检查当前节点是否在 codes 中或有子节点保留下来
		if codes[node.Value] || len(filteredChildren) > 0 {
			result = append(result, Node{
				Label:    node.Label,
				Value:    node.Value,
				Children: filteredChildren,
			})
		}
	}
	return result
}

// pullRoleDetail 获取角色详情
func pullRoleDetail(roleId int) (types.Role, error) {
	var r types.Role
	res := database.DB.Raw("select * from role where id = ?", roleId).Scan(&r)
	if res.RowsAffected == 0 {
		return r, fmt.Errorf("角色类型不存在")
	}
	if res.Error != nil {
		return r, res.Error
	}
	return r, nil
}

type buildTree struct {
	RoleId     int             `json:"role_id"`
	CasbinRole string          `json:"casbin_role"`
	Codes      map[string]bool `json:"codes"`
	Tree       []Node          `json:"tree"`
}

// getRoleId 获取角色id
func (buildTree *buildTree) getRoleId(ctx *gin.Context) error {
	accountDetail, err := account.GetUserInfo(ctx)
	if err != nil {
		return err
	}
	buildTree.RoleId = accountDetail.RoleId
	return nil
}

// getCasbinRole 获取casbin中的角色
func (buildTree *buildTree) getCasbinRole() error {
	roleDetail, err := pullRoleDetail(buildTree.RoleId)
	if err != nil {
		return err
	}
	buildTree.CasbinRole = roleDetail.CasbinRole
	return nil
}

// getCodes 获取权限codes
func (buildTree *buildTree) getCodes() error {
	filteredPolicy, err := casbins.E.GetRolesForUser(buildTree.CasbinRole)
	if err != nil {
		return err
	}
	buildTree.Codes = slice.ToMap(filteredPolicy)
	return nil
}

// buildTree 构建权限树
func (buildTree *buildTree) buildTree() {
	buildTree.Tree = filterTree(tree, buildTree.Codes)
	fmt.Println("tree===>")
	fmt.Println(buildTree.Tree)
	fmt.Println("tree<===")
}

// BuildUserTree 构建用户权限树
// 根据登录账号uid => 查询当前账号角色id => 根据角色id查询角色 => 通过角色获取权限codes => 根据codes过滤完整权限树 => 返回用户权限树
func BuildUserTree(ctx *gin.Context) ([]Node, error) {
	buildTree := buildTree{}
	err := buildTree.getRoleId(ctx)
	if err != nil {
		return buildTree.Tree, err
	}
	err = buildTree.getCasbinRole()
	if err != nil {
		return buildTree.Tree, err
	}
	err = buildTree.getCodes()
	if err != nil {
		return buildTree.Tree, err
	}
	buildTree.buildTree()
	return buildTree.Tree, nil
}
