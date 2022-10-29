package models

type User struct {
	BaseModel
	UserName string `json:"username" gorm:"type:varchar(100);comment:用户名"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
