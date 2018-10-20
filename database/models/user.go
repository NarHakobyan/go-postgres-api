package models

import (
	"time"

	"github.com/NarHakobyan/go-postgres-api/core/constants/roles"
	. "github.com/NarHakobyan/go-postgres-api/database"
	"github.com/fatih/structs"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//go:generate goqueryset -in user.go

// User struct represents user model.
// gen:qs
type User struct {
	Model
	Name     string         `form:"name" json:"name" valid:"required~Name is required"`
	Email    string         `form:"email" gorm:"unique_index" json:"email" valid:"email~Email isn't valid"`
	Stores   []Store        `form:"stores" json:"stores" valid:"email~Stores isn't valid"`
	Password string         `form:"password" json:"-" valid:"required~Password is required"`
	BirthDay time.Time      `form:"birthday" json:"birthday" valid:"required~Birth day is required" time_format:"02-01-2006"`
	Role     roles.RoleType `gorm:"default:0" json:"role"`
}

var UserQuery = NewUserQuerySet(Db)

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	scope.SetColumn("Password", hash)

	return nil
}

func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	cost, _ := bcrypt.Cost([]byte(u.Password))

	if cost == 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		scope.SetColumn("Password", hash)
	}
	return nil
}

func (u *User) ToJSON() map[string]interface{} {
	userObject := structs.Map(u)
	delete(userObject, "password")
	return userObject
}
func (u *User) ComparePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}

	return true
}
