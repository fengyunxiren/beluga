package forms

type RegisterForm struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,password"`
	Email    string `json:"email" binding:"required,email"`
}
