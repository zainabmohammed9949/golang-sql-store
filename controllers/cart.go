package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/zainabmohammed9949/golang-sql-store/models"
)

func AddToCart(DD *gorm.DB, res http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		http.ServeFile(res, req, "homepage.html")

	}

	sellerproduct := []models.Product{}
	txt10 := req.FormValue("product_name")
	resuls := DD.Where("product_name <>?", txt10).First(sellerproduct)

	ckeck := []models.ProductUser{}
	if ckeck == nil {
		DD.AutoMigrate(&models.ProductUser{})
		DD.NewRecord(&resuls)

	}
	DD.NewRecord(&resuls)

	log.Print(" added to cart ")
	json.NewEncoder(res).Encode(resuls)
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

}
func RemoveItem(DD *gorm.DB, res http.ResponseWriter, req *http.Request) {

	if req.Method != "DELETE" {
		http.ServeFile(res, req, "homepage.html")

	}
	txt1 := req.FormValue("product_name")

	userproduct := []models.ProductUser{}
	DD.Where("product_name=?", txt1).Delete(&userproduct)
	var _, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Printf("deleted from cart ")

}

func InstantBuy() {

	log.Printf(" deleted fromcart ")

}
