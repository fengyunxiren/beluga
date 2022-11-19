package models

type Role struct {
	BaseModel
	OrgId uint64 `json:"org_id"`
	Name  string `json:"name"`
	Code  string `json:"code" gorm:"uniqueIndex"`
	Desc  string `json:"desc"`
}

type RoleExclusion struct {
	BaseModel
	OrgId   uint64 `json:"org_id"`
	Role1Id uint64 `json:"role1_id"`
	Role2Id uint64 `json:"role2_id"`
}

type Permission struct {
	BaseModel
	OrgId    uint64 `json:"org_id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Resource string `json:"resource" gorm:"uniqueIndex"`
	Action   string `json:"action"`
}

type DataPermission struct {
	BaseModel
	OrgId    uint64 `json:"org_id"`
	PermId   uint64 `json:"perm_id"`
	DataKey  string `json:"data_key"`
	ParentId uint64 `json:"parent_id"`
}

type RolePermissionMap struct {
	BaseModel
	OrgId  uint64 `json:"org_id"`
	RoleId uint64 `json:"role_id"`
	PermId uint64 `json:"perm_id"`
}

type RoleDataPermission struct {
	BaseModel
	OrgId      uint64 `json:"org_id"`
	RoleId     uint64 `json:"role_id"`
	DataPermId uint64 `json:"data_perm_id"`
	DataValue  string `json:"data_value"`
}

type Menu struct {
	BaseModel
	OrgId    uint64 `json:"org_id"`
	Name     string `json:"name"`
	Code     string `json:"code" gorm:"uniqueIndex"`
	PermId   uint64 `json:"perm_id"`
	ParentId uint64 `json:"parent_id"`
	Desc     string `json:"desc"`
}

type Widget struct {
	BaseModel
	OrgId  uint64 `json:"org_id"`
	Name   string `json:"name"`
	Code   string `json:"code" gorm:"uniqueIndex"`
	PermId uint64 `json:"perm_id"`
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
