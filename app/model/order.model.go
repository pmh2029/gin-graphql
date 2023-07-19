package model

import "time"

// OrdersTableName TableName
var OrdersTableName = "orders"

type Order struct {
	ID          int        `gorm:"column:id;primaryKey;type:bigint;not null;autoIncrement" mapstructure:"id" json:"id"`
	PhoneNumber string     `gorm:"column:phone_number;type:varchar(11);not null" mapstructure:"phone_number" json:"phone_number"`
	Address     string     `gorm:"column:address;type:text;not null" mapstructure:"address" json:"address"`
	OrderTime   time.Time  `gorm:"column:order_time;type:timestamp;not null;default:current_timestamp" mapstructure:"order_time" json:"order_time"`
	Total       float64    `gorm:"column:total;type:numeric;not null" mapstructure:"total" json:"total"`
	Status      int        `gorm:"column:status;type:integer;not null;default:0" mapstructure:"status" json:"status"`
	UserID      int        `gorm:"column:user_id;type:bigint;not null" mapstructure:"user_id" json:"user_id"`
	User        User       `gorm:"foreignKey:UserID" json:"user"`
	Products    []*Product `gorm:"many2many:order_products" json:"products"`
	BaseModel
}

// TableName func
func (i *Order) TableName() string {
	return OrdersTableName
}
