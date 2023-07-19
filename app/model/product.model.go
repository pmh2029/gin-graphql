package model

// ProductsTableName TableName
var ProductsTableName = "products"

type Product struct {
	ID          int      `gorm:"column:id;primaryKey;type:bigint;not null;autoIncrement" mapstructure:"id" json:"id"`
	ProductName string   `gorm:"column:product_name;type:varchar(50);not null;uniqueIndex:product" mapstructure:"product_name" json:"product_name"`
	Description *string  `gorm:"column:description;type:text" mapstructure:"description" json:"description"`
	Price       float64  `gorm:"column:price;type:numeric;not null" mapstructure:"price" json:"price"`
	InStock     int      `gorm:"column:in_stock;type:integer;not null" mapstructure:"in_stock" json:"in_stock"`
	Url         string   `gorm:"column:url;type:text;not null" mapstructure:"url" json:"url"`
	BrandID     int      `gorm:"column:brand_id;type:bigint;not null;uniqueIndex:product" mapstructure:"brand_id" json:"brand_id"`
	Brand       Brand    `gorm:"foreignKey:BrandID" json:"brand"`
	Orders      []*Order `gorm:"many2many:order_products" json:"orders"`
	BaseModel
}

// TableName func
func (i *Product) TableName() string {
	return ProductsTableName
}
