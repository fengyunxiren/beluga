package models

type User struct {
	BaseModel
	UserName string `json:"username" gorm:"type:varchar(100);unique;not null;comment:用户名"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
