package storage

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopi.com/gopi/util"
)

// GopiDB db
var GopiDB *gorm.DB

// PrepareDB make new instance for gopi database
func PrepareDB() {
	var err error
	config := util.GetConfigInstance("database")
	username := config.GetString("gopi.username")
	password := config.GetString("gopi.password")
	dbName := config.GetString("gopi.name")
	mysqlConfig := fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		dbName,
	)

	GopiDB, err = gorm.Open("mysql", mysqlConfig)
	if err != nil {
		log.Fatal(err)
	}
}
