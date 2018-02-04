package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/narhakobyan/go-pg-api/database"
	"github.com/narhakobyan/go-pg-api/database/models"
)

type userController struct {
}

func (c *userController) GetUsers(context *gin.Context) {
	var users []models.User
	database.Db.Find(&users)
	context.JSON(http.StatusOK, users)
}

func (c *userController) GetUser(context *gin.Context) {
	var user models.User

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user id",
		})
		return
	}

	database.Db.Find(&user, id)
	context.JSON(http.StatusOK, user)
}

func (c *userController) PostUser(context *gin.Context) {
	var user models.User
	if err := context.BindJSON(&user); err != nil {
		context.Status(400)
		return
	}
	user.BirthDay = time.Now()
	database.Db.Create(&user)
	context.JSON(200, user)
}

var UserController = &userController{}
