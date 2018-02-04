package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/narhakobyan/go-pg-api/config"
	"github.com/narhakobyan/go-pg-api/http/router"
	"github.com/spf13/viper"
)

func main() {

	if viper.Get("env") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "WORKING")
	})
	router.Router.Run(":" + config.C.Server.Port)
}
