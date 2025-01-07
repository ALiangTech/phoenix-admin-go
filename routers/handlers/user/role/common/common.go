package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"phoenix-go-admin/database"
	"phoenix-go-admin/routers/handlers/user/account"
	casbins "phoenix-go-admin/routers/middlewares/casbin"
	"phoenix-go-admin/utils/slice"
)

type Role struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CasbinRole string `json:"casbin_role"`
	CreateAt   string `json:"create_at"`
	UpdateAt   string `json:"update_at"`
}

// GetRoleDetailByRoleId 根据role_id 查询角色信息
func GetRoleDetailByRoleId(roleId int) (Role, error) {
	var r Role
	res := database.DB.Raw("select * from role where id = ?", roleId).Scan(&r)
	if res.RowsAffected == 0 {
		return r, fmt.Errorf("角色类型不存在")
	}
	if res.Error != nil {
		return r, res.Error
	}
	return r, nil
}

// GetRoleDetailByContext 通过上下文ctx 获取当前用户角色信息
func GetRoleDetailByContext(ctx *gin.Context) (Role, error) {
	accountDetail, err := account.GetUserInfo(ctx)
	if err != nil {
		return Role{}, err
	}
	return GetRoleDetailByRoleId(accountDetail.RoleId)
}

// GetRoleCodes 根据角色查询权限码
func GetRoleCodes(casbinRole string) ([]string, error) {
	codes, err := casbins.E.GetRolesForUser(casbinRole)
	return codes, err
}

// ValidateCodes 校验codes的合法性
func ValidateCodes(codes []string, casbinRole string) (bool, error) {
	originCodes, err := GetRoleCodes(casbinRole)
	if err != nil {
		return false, err
	}
	// 判断codes中的值是否都存在于originCodes中
	for _, code := range codes {
		if !slice.Contains(originCodes, code) {
			return false, fmt.Errorf("存在不合法的权限码")
		}
	}
	return true, nil
}
