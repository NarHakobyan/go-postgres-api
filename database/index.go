package database

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var Db *gorm.DB

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

func init() {
	var err error

	govalidator.SetFieldsRequiredByDefault(false)

	Db, err = gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", viper.GetString("db.host"), viper.GetString("db.user"), viper.GetString("db.database"), viper.GetString("db.password")))
	if err != nil {
		panic("failed to connect database")
	}
}
