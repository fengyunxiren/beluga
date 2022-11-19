package models

type Group struct {
	BaseModel
	Name string `json:"name"`
	Code string `json:"code" gorm:"uniqueIndex"`
	Desc string `json:"desc"`
}

type GroupMember struct {
	BaseModel
	GroupId uint64 `json:"group_id"`
	UserId  uint64 `json:"user_id"`
}
