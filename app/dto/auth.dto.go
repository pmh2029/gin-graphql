package dto

type SignUpDto struct {
	Email    string `json:"email" binding:"required" mapstructure:"email"`
	Username string `json:"username" binding:"required" mapstructure:"username"`
	Password string `json:"password" binding:"required" mapstructure:"password"`
}

type SignInDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
