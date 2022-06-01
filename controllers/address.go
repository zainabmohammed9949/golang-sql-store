package controllers

import (
	"database/sql"
	"log"

	//"github.com/gin-gonic/gin"
	"github.com/zainabmohammed9949/golang-sql-store/models"
)

func AddAddress(Db *sql.DB, a models.Address, u models.User) error {

	rows, err := Db.Query("INSERT  INTO address(home,city) VALUES (?,?) WHERE email =?", a.House, a.City, u.Email)
	if err != nil {
		log.Println("Invalid")
		return err
	}
	log.Printf("%d address created", rows)
	return nil

}

func EditaAddress(Db sql.DB, a models.Address, u models.User) (string, error) {
	rows, err := Db.Query("UPDATE  address SET home=?,city=? WHERE email=? ", a.House, a.City, u.Email)
	if err != nil {
		return "invalid", err
	}
	log.Printf("%d address edited", rows)

	return "editing successflly", nil
}
func DeleteAddress(Db sql.DB, u models.User) (string, error) {
	rows, err := Db.Query("DELETE * FROM  address  WHERE email=? ", u.Email)
	if err != nil {
		return "NO ADDRESS TO DELETE", err
	}
	log.Printf("%d address edited", rows)

	return "deleted successflly", nil
}
