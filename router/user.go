package router

import (
	"github.com/gin-gonic/gin"
	"github.com/narhakobyan/go-pg-api/database"
	"github.com/narhakobyan/go-pg-api/database/models"
	"net/http"
	"strconv"
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

	UserRouter.GET("/:id", func(context *gin.Context) {
		var user models.User

		id, err := strconv.ParseInt(context.Param("id"), 0, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid user id",
			})
			return
		}

		database.Db.Find(&user, id)
		context.JSON(http.StatusOK, user)
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
