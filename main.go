package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/narhakobyan/go-pg-api/config"
	"github.com/narhakobyan/go-pg-api/router"
)

func main() {

	router.Router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "WORKING")
	})
	router.Router.Run(":" + config.C.Server.Port)
}
