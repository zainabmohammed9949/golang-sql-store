package database

import (
	"fmt"

	"github.com/zainabmohammed9949/golang-sql-store/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("zainab:0000@/shopdb"), &gorm.Config{})
	if err != nil {
		panic("لم نستطع الاتصال بقاعده البيانات ")
	}
	fmt.Println("database connected")

	//db.AutoMigrate(&models.Seller{})
	db.AutoMigrate(&models.User{}, &models.UserAddress{}, &models.Product{}, &models.Seller{}, &models.Order{})

	DB = db

}

func GetDB() *gorm.DB {
	var db *gorm.DB
	return db
}
