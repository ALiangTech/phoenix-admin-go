package types

type Role struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CasbinRole string `json:"casbin_role"`
	CreateAt   string `json:"create_at"`
	UpdateAt   string `json:"update_at"`
}
