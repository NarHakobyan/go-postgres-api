package database

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	Db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=go_rest_api sslmode=disable password=admin")
	if err != nil {
		panic("failed to connect database")
	}
}
