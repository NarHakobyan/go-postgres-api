package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/narhakobyan/go-pg-api/core/auth"
	"github.com/narhakobyan/go-pg-api/core/constants"
	"github.com/narhakobyan/go-pg-api/core/constants/roles"
	"github.com/narhakobyan/go-pg-api/database"
	"github.com/narhakobyan/go-pg-api/database/models"
)

func AuthMiddleware(roles []roles.RoleType) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenString, &auth.Claims{}, func(token2 *jwt.Token) (interface{}, error) {
			return auth.SigningKey, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*auth.Claims); ok && token.Valid {
			var user models.User
			database.Db.Find(&user, claims.UserId)
			c.Set(constants.AuthUser, &user)
			for _, role := range roles {
				if role == user.Role {
					c.Next()
					return
				}
			}
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Status(http.StatusUnauthorized)

	}
}
