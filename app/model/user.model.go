package model

// UsersTableName TableName
var UsersTableName = "users"

type User struct {
	ID       int      `gorm:"column:id;primaryKey;type:bigint;not null;autoIncrement" mapstructure:"id" json:"id"`
	Username string   `gorm:"column:username;type:varchar(50);not null;unique" mapstructure:"username" json:"username"`
	Email    string   `gorm:"column:email;type:varchar(50);not null;unique" mapstructure:"email" json:"email"`
	Password string   `gorm:"column:password;not null" mapstructure:"password" json:"password"`
	IsAdmin  bool     `gorm:"column:is_admin;type:boolean;default:false" mapstructure:"is_admin" json:"is_admin"`
	Orders   []*Order `gorm:"foreignKey:UserID" json:"orders"`
	BaseModel
}

// TableName func
func (i *User) TableName() string {
	return UsersTableName
}
