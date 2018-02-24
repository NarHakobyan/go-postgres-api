package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/narhakobyan/go-pg-api/core/constants/roles"
	. "github.com/narhakobyan/go-pg-api/database"
	. "github.com/narhakobyan/go-pg-api/database/models"
)

type userController struct{}

func (controller *userController) GetUsers(context *gin.Context) {
	var users []User
	Db.Find(&users)
	context.JSON(http.StatusOK, users)
}

func (controller *userController) GetUser(context *gin.Context) {
	var user User
	var id int
	var err error
	if id, err = strconv.Atoi(context.Param("id")); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user id",
		})
		return
	}
	Db.Find(&user, id)

	if structs.IsZero(user) {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	context.JSON(http.StatusOK, user)
}

func (controller *userController) UpdateUser(context *gin.Context) {
	var user User
	var userBody User

	if err := context.BindJSON(&userBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Data is not valid",
		})
		return
	}
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id",
		})
		return
	}

	Db.Find(&user, id)

	if structs.IsZero(user) {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	Db.Model(&user).Updates(userBody)

	context.JSON(http.StatusOK, user)
}

func (controller *userController) PostUser(context *gin.Context) {
	var user User

	if err := context.ShouldBindWith(&user, binding.FormPost); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	user.Role = roles.UserRole
	if valid, err := govalidator.ValidateStruct(user); err != nil || valid == false {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": strings.Split(err.Error(), ";")})
	}

	Db.Create(&user)
	context.JSON(http.StatusOK, gin.H{"data": user})
}

var UserController = &userController{}
