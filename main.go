package main

import (
	"github.com/narhakobyan/go-pg-api/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/narhakobyan/go-pg-api/database"
)

func main() {

	defer database.Db.Close();

	router.Router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "WORKING")
	})
	router.Router.Run(":8000")
}
