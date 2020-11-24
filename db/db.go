package db

import (
	//"encoding/json"
	"go-stock/model"

	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	//"github.com/jinzhu/gorm"
)

var db *gorm.DB
var server = "192.168.0.80"
var port = 1433
var user = "sa"
var password = "123456"
var dbname = "UMS_Mobile"

func GetDB() *gorm.DB {
	return db
}

func SetupDB() {

	dsn := "sqlserver://sa:123456@192.168.0.80:1433?database=UMS_Mobile"
	database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to create connection Error: " + err.Error())
	}

	database.AutoMigrate(&model.User{}, &model.Product{}, &model.Transaction{})
	db = database

}
