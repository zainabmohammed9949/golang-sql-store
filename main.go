package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zainabmohammed9949/golang-sql-store/routes"
	//"github.com/zainabmohammed9949/golang-sql-store/tokens"
)

func main() {
	fmt.Println("Gorm API PROJECT STARTED")
	r := mux.NewRouter()
	routes.UserRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
