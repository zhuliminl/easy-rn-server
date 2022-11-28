package dto

type User struct {
	UserId   string `json:"user_id"`
	Username string `json:"username" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
}

type UserRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UserRegisterByEmail struct {
	Username string `json:"username"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterByPhone struct {
	Username string `json:"username"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginByEmail struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginByPhone struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}
