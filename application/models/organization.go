package models

type Organization struct {
	BaseModel
	Name string `json:"name"`
	Code string `json:"code" gorm:"uniqueIndex"`
	Desc string `json:"desc"`
}

type OrganizationMember struct {
	BaseModel
	OrgId     uint64 `json:"org_id"`
	UserId    uint64 `json:"user_id"`
	OrgNumber string `json:"org_number" gorm:"uniqueIndex"`
}

type Department struct {
	BaseModel
	OrgId    uint64 `json:"org_id"`
	Name     string `json:"name"`
	Code     string `json:"code" gorm:"uniqueIndex"`
	Desc     string `json:"desc"`
	ParentId uint32 `json:"parent_id" gorm:"default:0"`
}

type DepartmentMember struct {
	BaseModel
	OrgId     uint64 `json:"org_id"`
	DepId     uint64 `json:"dep_id"`
	OrgNumber string `json:"org_number"`
}

type OrganizationGroup struct {
	BaseModel
	OrgId uint64 `json:"org_id"`
	Name  string `json:"name"`
	Code  string `json:"code" gorm:"uniqueIndex"`
	Desc  string `json:"desc"`
}

type OrganizationGroupMember struct {
	BaseModel
	OrgId     uint64 `json:"org_id"`
	DepId     uint64 `json:"dep_id"`
	OrgNumber string `json:"org_number"`
}

type OrganizationRank struct {
	BaseModel
	OrgId    uint64 `json:"org_id"`
	Name     string `json:"name"`
	Code     string `json:"code" gorm:"uniqueIndex"`
	Desc     string `json:"desc"`
	ParentId uint64 `json:"parent_id" gorm:"default:0"`
}

type OrganizationRankLevel struct {
	BaseModel
	OrgId     uint64 `json:"org_id"`
	RankId    uint64 `json:"rank_id"`
	Level     uint64 `json:"level"`
	LevelName uint64 `json:"level_name"`
}

type OrganizationMemberRank struct {
	BaseModel
	OrgId     uint64 `json:"org_id"`
	LevelId   uint64 `json:"level_id"`
	OrgNumber string `json:"org_number"`
}

type OrganizationMemberRole struct {
	BaseModel
	OrgId     uint64 `json:"org_id"`
	OrgNumber string `json:"org_number"`
	RoleId    uint64 `json:"role_id"`
}

type OrganizationDepartmentRole struct {
	BaseModel
	OrgId  uint64 `json:"org_id"`
	DepId  uint64 `json:"dep_id"`
	RoleId uint64 `json:"role_id"`
}

type OrganizationGroupRole struct {
	BaseModel
	OrgId   uint64 `json:"org_id"`
	GroupId uint64 `json:"group_id"`
	RoleId  uint64 `json:"role_id"`
}

type OrganizationRankLevelRole struct {
	BaseModel
	OrgId   uint64 `json:"org_id"`
	LevelId uint64 `json:"level_id"`
	RoleId  uint64 `json:"role_id"`
}
