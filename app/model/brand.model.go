package model

// BrandsTableName TableName
var BrandsTableName = "brands"

type Brand struct {
	ID        int       `gorm:"column:id;primaryKey;type:bigint;not null;autoIncrement" mapstructure:"id" json:"id"`
	BrandName string    `gorm:"column:brand_name;type:varchar(50);not null;unique" mapstructure:"brand_name" json:"brand_name"`
	Products  []Product `gorm:"foreignKey:BrandID" json:"products"`
	BaseModel
}

// TableName func
func (i *Brand) TableName() string {
	return BrandsTableName
}
