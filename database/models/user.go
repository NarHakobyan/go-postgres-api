package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type (
	UserModel struct {
		gorm.Model
		Name     string    `json:"name"`
		Password int       `json:"password"`
		BirthDay time.Time `json:"birth_day"`
	}
)
