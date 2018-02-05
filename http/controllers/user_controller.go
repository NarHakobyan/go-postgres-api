package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
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

	if user.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}
	context.JSON(http.StatusOK, user)
}

func (c *userController) UpdateUser(context *gin.Context) {
	var user models.User
	var userBody models.User

	if err := context.BindJSON(&userBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Data is not valid",
		})
		return
	}

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user id",
		})
		return
	}

	database.Db.Find(&user, id)

	if user.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}
	database.Db.Model(&user).Updates(userBody)

	context.JSON(http.StatusOK, user)
}

func (c *userController) PostUser(context *gin.Context) {
	var user models.User
	if err := context.BindJSON(&user); err != nil {
		context.Status(http.StatusBadRequest)
		return
	}
	user.BirthDay = time.Now()
	valid, err := govalidator.ValidateStruct(user)
	if valid == false {
		context.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})
		return
	}
	database.Db.Create(&user)
	context.JSON(200, user)
}

var UserController = &userController{}
