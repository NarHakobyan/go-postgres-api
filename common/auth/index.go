package auth

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

var SigningKey = []byte("test test")
