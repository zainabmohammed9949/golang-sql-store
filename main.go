package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zainabmohammed9949/golang-sql-store/routes"

	//"github.com/zainabmohammed9949/golang-sql-store/tokens"
	"github.com/zainabmohammed9949/golang-sql-store/database"
)

func main() {
	fmt.Println("Gorm API PROJECT STARTED")
	r := mux.NewRouter()
	routes.UserRoutes(r)
	database.Init()
	http.Handle("/", r)
	fmt.Println("tables created")
	log.Fatal(http.ListenAndServe("localhost:3306", r))
}
