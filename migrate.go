package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/narhakobyan/go-pg-api/database"
	"github.com/narhakobyan/go-pg-api/database/models"
)

func main() {
	database.Db.AutoMigrate(&models.User{})
}
