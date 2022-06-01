package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
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
func signup(DD *gorm.DB, res http.ResponseWriter, req *http.Request) error {
	if req.Method != "POST" {
		http.ServeFile(res, req, "/signup.html")
		return nil
	}
	email := req.FormValue("email")
	password := req.FormValue("password")
	user_Name := req.FormValue("u_name")
	phone := req.FormValue("phone")

	user := models.User{User_Name: &user_Name, Password: &password, Phone: &phone, Email: &email}
	results := DD.Create(&user)
	err := results.Error
	return err
}
func searchProduct(DB *gorm.DB, res http.ResponseWriter, r *http.Request, p models.Product) (*gorm.DB, error) {
	if r.Method != "POST" {
		http.ServeFile(res, r, "homepage.html")

	}
	name := r.FormValue("product_name")
	prod := DB.First(&p, "product_name=?", name)
	if prod == nil {
		panic("there are no product in this name")
		http.Redirect(res, r, "/homepage", 302)
	}

	return prod, nil

}
func insertProduct(ctx context.Context, DD *gorm.DB, res http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(res, r, "seller_profile.html")

	}
	txt1 := r.FormValue("product_name")
	txt2 := r.FormValue("Price")
	txt3 := r.FormValue("Image")
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	product := models.Product{Product_Name: &txt1, Price: &txt2, Image: &txt3}
	DD.NewRecord(product)
	results := DD.Create(&product)

	err := results.Error

	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
	}
	log.Printf("%d products created ", "{}", results)
	json.NewEncoder(res).Encode(results)
}
func deleteProduct(ctx context.Context, DD *gorm.DB, res http.ResponseWriter, req *http.Request) {
	if req.Method != "Delete" {
		http.ServeFile(res, req, "seller_profile.html")

	}
	product := []models.Product{}
	txt1 := req.FormValue("product_name")
	prod := DD.Where("product_name <>?", txt1).First(product)
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	results := DD.Delete(&prod)

	err := results.Error

	if err != nil {
		log.Printf("Error %s when deleting row into products table", err)
	}
	log.Printf("%d product deleted", results)
	json.NewEncoder(res).Encode(results)
}

func deleteAccount(ctx context.Context, DD *gorm.DB, res http.ResponseWriter, req *http.Request) (*http.Response, error) {
	if req.Method != "Delete" {
		http.ServeFile(res, req, "seller_profile.html")

	}
	acc := []models.Seller{}
	txt10 := req.FormValue("product_name")
	resul := DD.Where("product_name <>?", txt10).First(acc)
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	results := DD.Delete(&resul)

	err := results.Error

	if err != nil {
		log.Printf("Error %s when deleting this seller table", err)
		return &http.Response{Status: "Account Not Deleted"}, err
	}
	log.Printf("%d Account  deleted", results)
	return &http.Response{Status: "Deleted"}, nil
}
