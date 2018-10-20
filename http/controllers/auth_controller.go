package controllers

import (
	"time"

	"github.com/NarHakobyan/go-postgres-api/core/auth"
	. "github.com/NarHakobyan/go-postgres-api/database/models"
	"github.com/NarHakobyan/go-postgres-api/http/response"
	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type authController struct{}

type Login struct {
	Email    string `form:"email" json:"email" valid:"required~Email is required, email~Email isn't valid"`
	Password string `form:"password" json:"password" valid:"required~Password is required"`
}

func (controller *authController) PostLogin(context *response.Context) {
	var login Login
	var user User

	if err := context.ShouldBindWith(&login, binding.FormPost); err != nil {
		context.BadRequest(err.Error(), nil)
		return
	}

	if _, err := govalidator.ValidateStruct(login); err != nil {
		context.UnprocessableEntity("", err.(govalidator.Errors).Errors())
		return
	}

	UserQuery.EmailEq(login.Email).One(&user)

	if empty := structs.IsZero(user); empty == true {
		context.BadRequest("Email or password is incorrect", nil)
		return
	}

	claims := auth.Claims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
			Issuer:    "test",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token, err := tokenClaims.SignedString(auth.SigningKey); err == nil {
		context.Ok("Successfully loggedIn", gin.H{
			"user":  user,
			"token": token,
		})
	} else {
		context.InternalServerError("", err.Error())
	}
}

func (controller *authController) PostRegister(context *response.Context) {

}

func (controller *authController) GetMyProfile(context *response.Context) {
	user, _ := context.GetUser()

	context.Ok("", user)
}

var AuthController = &authController{}
