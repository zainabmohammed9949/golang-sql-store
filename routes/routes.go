package routes

import (
	"github.com/gorilla/mux"
	"github.com/zainabmohammed9949/golang-sql-store/controllers"
)

var UserRoutes = func(router *mux.Router) {
	router.HandleFunc("/users/signup", controllers.Signup()).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login()).Methods("POST")
	router.HandleFunc("/users/productview", controllers.SearchProduct()).Methods("GET")
	router.HandleFunc("/users/search", controllers.SearchProductByQuery()).Methods("GET")
	router.HandleFunc("/admin/addproduct", controllers.ProduclViewerAdmins()).Methods("Post")
}
