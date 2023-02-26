package config

import (
	"github.com/akhilesh-ingle-ge/crud-web-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dsn = "postgres://postgres:12345@localhost:5432/test?sslmode=disable"

func SetUpDB() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}
	db.AutoMigrate(&models.Student{})
	DB = db
}
