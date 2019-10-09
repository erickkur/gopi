package migration

import (
	"github.com/jinzhu/gorm"
	"gopi.com/gopi/storage"
	gormigrate "gopkg.in/gormigrate.v1"
)

var Migration20181113233300 = &gormigrate.Migration{
	ID: "20181113233300",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&storage.User{}).Error
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.DropTable("users").Error
	},
}
