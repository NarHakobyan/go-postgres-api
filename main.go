package main

import (
	"flag"
	"fmt"

	_ "github.com/NarHakobyan/go-postgres-api/config"
	"github.com/NarHakobyan/go-postgres-api/database"
	"github.com/NarHakobyan/go-postgres-api/database/models"
	. "github.com/NarHakobyan/go-postgres-api/http/router"
	"github.com/gin-gonic/gin"
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

	Router.Run(":" + viper.GetString("server.port"))
}
