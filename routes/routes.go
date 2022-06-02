package routes

import (
	//"log"
	//"net/http"

	"github.com/gorilla/mux"
	"github.com/zainabmohammed9949/golang-sql-store/controllers"
)

var UserRoutes = func(router *mux.Router) {
	router.HandleFunc("/users/signup/{name}", controllers.Signup)
	//router.HandleFunc("/users/orders", controllers.CompleteOrder()).Methods("PUT")
	//router.HandleFunc("/users/profile/{user_name}", controllers.login()).Methods("POST")
	router.HandleFunc("/users/productview/{product_id}", controllers.SearchProduct)
	//router.HandleFunc("/user/profile", controllers.AddUserAddress()).Methods("PUT")

	prodroute := router.PathPrefix("/products").Subrouter()
	prodroute.HandleFunc("/", controllers.AllProducts)
	prodroute.HandleFunc("/{product_name}", controllers.InsertProduct)
	//router.HandleFunc("/seller/signup", controllers.SellerSignup()).Methods("POST")
	//router.HandleFunc("seller/login", controllers.Login()).Methods("POST")
	router.HandleFunc("/Seller/profile/", controllers.DeleteAccount)
	router.HandleFunc("/deleted", controllers.DeleteProduct)

}
