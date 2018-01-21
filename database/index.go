package database

import (
	"github.com/jinzhu/gorm"
	"github.com/narhakobyan/go-pg-api/database/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

func init() {
	Db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=go_rest_api sslmode=disable password=admin")
	if err != nil {
		panic("failed to connect database")
	}
	Db.AutoMigrate(&models.UserModel{})
}
