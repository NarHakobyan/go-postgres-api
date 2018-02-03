package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/narhakobyan/go-pg-api/database"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	database.Model
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	BirthDay time.Time `json:"birth_day"`
	Role     int       `gorm:"default:1" json:"role"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	scope.SetColumn("Password", hash)

	return nil
}
func (user *User) ComparePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false
	}

	return true
}
