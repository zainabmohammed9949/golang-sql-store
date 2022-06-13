package main

import (
	//"fmt"
	//"log"
	//"net/http"
	//"github.com/jinzhu/gorm"

	//"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zainabmohammed9949/golang-sql-store/routes"

	//"github.com/zainabmohammed9949/golang-sql-store/tokens"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/zainabmohammed9949/golang-sql-store/database"
)

//var frontend embed.FS

func main() {

	//fmt.Println("Gorm API PROJECT STARTED")
	database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.SetUp(app)
	app.Listen(":8080")

	//r := mux.NewRouter()
	//database.Init()
	//http.Handle("/", r)
	//fmt.Println("tables created")
	//log.Fatal(http.ListenAndServe("localhost:3306", r))

	//stripped, err := fs.Sub(frontend, "frontend/dist")
	//if err != null {
	//    log.Fatalln(err)
	//}

	//frontendFS := http.FileServer(http.FS(stripped))
	//http.Handle("/", frontendFS)

	// ...
}
