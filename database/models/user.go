package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	gorm.Model
	Name     string    `json:"name"`
	Password string    `json:"password"`
	BirthDay time.Time `json:"birth_day"`
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
