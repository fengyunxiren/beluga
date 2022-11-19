package models

type Application struct {
	BaseModel
	OrgId     uint64 `json:"org_id"`
	AppKey    string `json:"app_key" gorm:"uniqueIndex"`
	AppSecret string `json:"app_secret"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
}

type ApplicationRole struct {
	BaseModel
	OrgId  uint64 `json:"org_id"`
	AppId  uint64 `json:"app_id"`
	RoleId uint64 `json:"role_id"`
}
