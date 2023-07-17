package dto

type UpdateUserDto struct {
	Email    *string `json:"email" mapstructure:"email" binding:"omitempty,min=1"`
	Username *string `json:"username" mapstructure:"username" binding:"omitempty,min=1"`
	Password *string `json:"password"  mapstructure:"password" binding:"omitempty,min=1"`
}
