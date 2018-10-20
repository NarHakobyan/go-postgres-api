package controllers

import (
	"strconv"
	"time"

	"github.com/NarHakobyan/go-postgres-api/core/constants/roles"
	. "github.com/NarHakobyan/go-postgres-api/database"
	. "github.com/NarHakobyan/go-postgres-api/database/models"
	"github.com/NarHakobyan/go-postgres-api/http/response"
	"github.com/asaskevich/govalidator"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin/binding"
)

type updateUser struct {
	Name     string    `form:"name"`
	Email    string    `form:"email"`
	Password string    `form:"password"`
	BirthDay time.Time `form:"birthday" time_format:"02-01-2006"`
}
type newUser struct {
	Name     string    `form:"name" valid:"required~Name is required"`
	Email    string    `form:"email" valid:"email~Email isn't valid"`
	Password string    `form:"password" valid:"required~Password is required"`
	BirthDay time.Time `form:"birthday" valid:"required~Birth day is required" time_format:"02-01-2006"`
	Role     roles.RoleType
}

type userController struct{}

func (controller *userController) GetUsers(context *response.Context) {
	var users []User
	Db.Find(&users)
	context.Ok("", users)
}

func (controller *userController) GetUser(context *response.Context) {
	var user User
	var id int
	var err error
	if id, err = strconv.Atoi(context.Param("id")); err != nil {
		context.BadRequest("Invalid user id", nil)
		return
	}
	Db.Find(&user, id)

	if structs.IsZero(user) {
		context.NotFound("User not found", nil)
		return
	}
	context.Ok("", user)
}

func (controller *userController) UpdateUser(context *response.Context) {
	var user User
	var userBody updateUser

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.BadRequest("Invalid user id", nil)
		return
	}

	if err := context.ShouldBindWith(&userBody, binding.FormMultipart); err != nil {
		context.BadRequest(err.Error(), nil)
		return
	}

	if _, err := govalidator.ValidateStruct(userBody); err != nil {
		context.UnprocessableEntity("", err.(govalidator.Errors).Errors())
		return
	}

	Db.First(&user, id)

	if structs.IsZero(user) {
		context.NotFound("User not found", nil)
		return
	}
	Db.Model(&user).Updates(userBody)

	context.Ok("User successfully updated", user)
}

func (controller *userController) PostUser(context *response.Context) {
	var userBody newUser

	if err := context.ShouldBindWith(&userBody, binding.FormMultipart); err != nil {
		context.BadRequest(err.Error(), nil)
		return
	}

	userBody.Role = roles.UserRole

	if _, err := govalidator.ValidateStruct(userBody); err != nil {
		context.UnprocessableEntity("", err.(govalidator.Errors).Errors())
		return
	}

	Db.Create(&userBody)

	context.Ok("User Successfully created", userBody)
}

var UserController = &userController{}
