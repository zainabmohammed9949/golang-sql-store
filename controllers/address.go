package controllers

import (
	"database/sql"
	"log"

	"net/http"

	"github.com/jinzhu/gorm"

	//"github.com/gin-gonic/gin"
	"github.com/zainabmohammed9949/eco-go/models"
	"github.com/zainabmohammed9949/golang-sql-store/models"
)

func AddUserAddress(DD *gorm.DB, res http.ResponseWriter, req *http.Request) error {
	home := req.FormValue("home_name")
	city := req.FormValue("city_name")

	address := models.UserAddress{House: &home, City: &city}
	results := DD.Create(&address)
	err := results.Error
	return err

	log.Printf("%d address created", results)
	return nil

}
func AddSellerAddress(DD *gorm.DB, res http.ResponseWriter, req *http.Request) error {
	store := req.FormValue("store_name")
	city := req.FormValue("city_name")

	address := models.SellerAddress{Store: &store, City: &city}
	results := DD.Create(&address)
	err := results.Error
	return err

	log.Printf("%d address created", results)
	return nil

}

func EditaAddress(Db sql.DB, DD *gorm.DB, res http.ResponseWriter, req *http.Request) {
	home := req.FormValue("home_name")
	city := req.FormValue("city_name")

	address := models.UserAddress{House: &home, City: &city}
	results := DD.Update(&address)

	log.Printf("%d address edited", results)

}
func DeleteAddress(Db sql.DB, u models.User) (string, error) {
	rows, err := Db.Query("DELETE * FROM  address  WHERE email=? ", u.Email)
	if err != nil {
		return "NO ADDRESS TO DELETE", err
	}
	log.Printf("%d address edited", rows)

	return "deleted successflly", nil
}
