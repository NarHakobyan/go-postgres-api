package main

import (
	"github.com/narhakobyan/go-pg-api/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router.Router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "WORKING")
	})
	router.Router.Run(":8000")
}
