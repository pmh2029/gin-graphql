package model

// OrderProductsTableName TableName
var OrderProductsTableName = "order_products"

type OrderProduct struct {
	ID        int `gorm:"column:id;primaryKey;type:bigint;not null;autoIncrement" mapstructure:"id" json:"id"`
	OrderID   int `gorm:"column:order_id;type:bigint;not null" mapstructure:"order_id" json:"order_id"`
	ProductID int `gorm:"column:product_id;type:bigint;not null" mapstructure:"product_id" json:"product_id"`
	Quantity  int `gorm:"column:quantity;type:int;not null" mapstructure:"quantity" json:"quantity"`
	BaseModel
}

// TableName func
func (i *OrderProduct) TableName() string {
	return OrderProductsTableName
}
