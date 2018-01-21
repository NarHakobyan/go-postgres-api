package main

import (
	"github.com/narhakobyan/go-pg-api/database/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/narhakobyan/go-pg-api/database"
)

func main() {
	database.Db.AutoMigrate(&models.UserModel{})
}
