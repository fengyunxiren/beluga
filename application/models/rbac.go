package models

type Role struct {
	BaseModel
	OrgId uint64 `json:"org_id"`
	Name  string `json:"name"`
	Code  string `json:"code" gorm:"uniqueIndex"`
	Desc  string `json:"desc"`
}

type Permission struct {
	BaseModel
	OrgId     uint64 `json:"org_id"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	Resource  string `json:"resource" gorm:"uniqueIndex"`
	Action    string `json:"action"`
	ParentId  uint64 `json:"parent_id"`
	SelectKey string `json:"select_key"`
	SelectUrl string `json:"select_url"`
}

type RolePermissionMap struct {
	BaseModel
	OrgId          uint64 `json:"org_id"`
	RoleId         uint64 `json:"role_id"`
	PermId         uint64 `json:"perm_id"`
	PermissionType uint64 `json:"perm_type"` // default, select, all
}

type RolePermissionSelect struct {
	BaseModel
	OrgId     uint64 `json:"org_id"`
	MPId      uint64 `json:"rp_id"`
	DataValue string `json:"data_value"`
}

type Menu struct {
	BaseModel
	OrgId    uint64 `json:"org_id"`
	Name     string `json:"name"`
	Code     string `json:"code" gorm:"uniqueIndex"`
	ParentId uint64 `json:"parent_id"`
	Desc     string `json:"desc"`
}

type Widget struct {
	BaseModel
	OrgId  uint64 `json:"org_id"`
	Name   string `json:"name"`
	Code   string `json:"code" gorm:"uniqueIndex"`
	MenuId uint64 `json:"menu_id"`
	Desc   string `json:"desc"`
}

type RoleMenuMap struct {
	BaseModel
	OrgId  uint64 `json:"org_id"`
	RoleId uint64 `json:"role_id"`
	MenuId uint64 `json:"menu_id"`
}

type RoleWidgetMap struct {
	BaseModel
	OrgId    uint64 `json:"org_id"`
	RoleId   uint64 `json:"role_id"`
	WidgetId uint64 `json:"widget_id"`
}
