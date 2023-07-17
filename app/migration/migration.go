package migration

import (
	"gin-graphql/app/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(model.User{})
}
