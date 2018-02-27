package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/narhakobyan/go-pg-api/config"
	"github.com/narhakobyan/go-pg-api/database"
	"github.com/narhakobyan/go-pg-api/database/models"
	. "github.com/narhakobyan/go-pg-api/http/router"
	"github.com/spf13/viper"
)

func main() {

	dbSync := flag.Bool("sdb", false, "Force sync database models")
	flag.Parse()

	if *dbSync == true {
		database.Db.AutoMigrate(&models.User{}, &models.Store{})
		fmt.Println("Database sync was successfully done")
		return
	}

	InitBaseRouter()

	if viper.Get("env") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())

	Router.Run(":" + config.C.Server.Port)
}
