package database

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/narhakobyan/go-pg-api/config"
)

var Db *gorm.DB

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type Where map[string]interface{}

func init() {
	var err error

	govalidator.SetFieldsRequiredByDefault(false)

	Db, err = gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", config.C.Database.Host, config.C.Database.User, config.C.Database.DbName, config.C.Database.Password))
	if err != nil {
		panic("failed to connect database")
	}
}
