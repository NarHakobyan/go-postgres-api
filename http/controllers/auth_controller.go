package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/narhakobyan/go-pg-api/common/auth"
	. "github.com/narhakobyan/go-pg-api/database"
	. "github.com/narhakobyan/go-pg-api/database/models"
)

type authController struct{}

type Login struct {
	Email    string `form:"email" json:"email" valid:"required~Email is required, email~Email isn't valid"`
	Password string `form:"password" json:"password" valid:"required~Password is required"`
}

func (c *authController) PostLogin(context *gin.Context) {
	var login Login
	var user User
	if err := context.ShouldBindWith(&login, binding.FormPost); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	if valid, err := govalidator.ValidateStruct(login); err != nil || valid == false {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": strings.Split(err.Error(), ";")})
	}
	Db.First(&user, map[string]string{"email": user.Email})
	if empty := structs.IsZero(user); empty == true {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Email or password is incorrect",
		})
		return
	}

	// Create the Claims
	claims := auth.Claims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if ss, err := token.SignedString(auth.SigningKey); err == nil {
		context.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"user":  user,
				"token": ss,
			},
		})
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}

func (c *authController) PostRegister(context *gin.Context) {

}

var AuthController = &authController{}
