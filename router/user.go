package router

import (
	"github.com/gin-gonic/gin"
	"github.com/narhakobyan/go-pg-api/database"
	"github.com/narhakobyan/go-pg-api/database/models"
	"net/http"
	"time"
)

var UserRouter *gin.RouterGroup

func init() {
	UserRouter = ApiRouter.Group("/users")

	UserRouter.GET("/", func(context *gin.Context) {
		var users []models.User
		database.Db.Find(&users)
		context.JSON(http.StatusOK, users)
	})

	UserRouter.POST("/", func(context *gin.Context) {
		var user models.User
		if err := context.BindJSON(&user); err != nil {
			context.Status(400)
			return
		}
		user.BirthDay = time.Now()
		database.Db.Create(&user)
		context.JSON(200, user)
	})
}
