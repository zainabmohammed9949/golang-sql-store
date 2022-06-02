package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/zainabmohammed9949/golang-sql-store/database"
	"github.com/zainabmohammed9949/golang-sql-store/models"

	"context"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerfyPassward(userPassword string, givenPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	valid := true
	msg := ""
	if err != nil {
		msg = "Login or Password is not correct"
		valid = false
	}
	return valid, msg
}

var dhh []string = []string{}

func Signup(res http.ResponseWriter, req *http.Request) {
	//res.Header().Set("content-type", "application/json")

	vars := mux.Vars(req)["name"]
	dhh = append(dhh, vars)
	email := req.FormValue("email")
	password := req.FormValue("password")
	user_Name := req.FormValue("name")
	phone := req.FormValue("phone")

	user := models.User{User_Name: &user_Name, Password: &password, Phone: &phone, Email: &email}
	dd := database.GetDB()
	dd.Create(&user)

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(dhh)

}
func SearchProduct(res http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(res, r, "homepage.html")

	}
	name := r.FormValue("product_name")
	DB := database.GetDB()
	product := models.Product{}
	prod := DB.First(&product, "product_name=?", name)
	if prod == nil {
		panic("there are no product in this name")
		//http.Redirect(res, r, "/homepage", 302)
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(prod)
}

func AllProducts(res http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(res, r, "seller_profile.html")

	}
	product := models.Product{}
	DD := database.GetDB()
	results := DD.First(&product)

	err := results.Error

	if err != nil {
		log.Printf("Error %s no products table", err)
	}
	log.Printf(" products showed ")
	json.NewEncoder(res).Encode(results)
}

func InsertProduct(res http.ResponseWriter, r *http.Request) {
	txt1 := r.FormValue("product_name")
	txt2 := r.FormValue("Price")
	txt3 := r.FormValue("Image")
	_, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	product := models.Product{Product_Name: &txt1, Price: &txt2, Image: &txt3}
	DD := database.GetDB()
	DD.NewRecord(product)
	results := DD.Create(&product)

	err := results.Error

	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
	}
	log.Printf(" products created ")
	json.NewEncoder(res).Encode(results)
}
func DeleteProduct(res http.ResponseWriter, req *http.Request) {

	product := []models.Product{}
	txt1 := req.FormValue("product_name")
	DD := database.GetDB()

	prod := DD.Where("product_name <>?", txt1).First(product)
	_, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	results := DD.Delete(&prod)

	err := results.Error

	if err != nil {
		log.Printf("Error %s when deleting row into products table", err)
	}
	log.Printf("product deleted")
	json.NewEncoder(res).Encode(results)
}

func DeleteAccount(res http.ResponseWriter, req *http.Request) {

	acc := []models.Seller{}
	txt10 := req.FormValue("product_name")
	DD := database.GetDB()

	resul := DD.Where("product_name <>?", txt10).First(acc)
	_, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	results := DD.Delete(&resul)

	err := results.Error

	if err != nil {
		log.Printf("Error %s when deleting this seller table", err)

	}
	log.Printf("Account  deleted")

}
