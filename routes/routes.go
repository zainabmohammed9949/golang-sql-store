package routes

import (
	"github.com/gorilla/mux"
	"github.com/zainabmohammed9949/golang-sql-store/controllers"
)

var UserRoutes = func(router *mux.Router) {
	router.HandleFunc("/users/signup", controllers.Signup()).Methods("POST")
	router.HandleFunc("/users/orders", controllers.CompleteOrder()).Methods("PUT")
	router.HandleFunc("/users/profile/{user_name}", controllers.login()).Methods("POST")
	router.HandleFunc("/users/productview/{product_id}", controllers.SearchProduct()).Methods("GET")
	router.HandleFunc("/user/profile", controllers.UserProfile()).Methods("POST")

	prodroute := router.PathPrefix("/products").Subrouter()
	prodroute.HandleFunc("/", controllers.AllProducts())
	prodroute.HandleFunc("/{product_name}", controllers.ShowProducts())
	router.HandleFunc("/seller/signup", controllers.SellerSignup()).Methods("POST")
	router.HandleFunc("seller/login", controllers.Login()).Methods("POST")
	router.HandleFunc("/Seller/profile/{seller_name}", controllers.NewSeller()).Methods("POST")
	router.HandleFunc("/seller/profile/allproducts", controllers.DeleteProduct()).Methods("GET")
	//log.Fatal(http.ListenAndServe(":8081,router"))

}
