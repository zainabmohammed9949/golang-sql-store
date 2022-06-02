package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zainabmohammed9949/golang-sql-store/models"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "zainab:0000@tcp(localhost:3306)/shopdb")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
func Init() {
	Connect()
	db = GetDB()
	db.AutoMigrate(&models.Seller{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Product{})

}
