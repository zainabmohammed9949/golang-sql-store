package main

import (
	"log"
	"os"

	"github.com/zainabmohammed9949/eco-go/controllers"
	"github.com/zainabmohammed9949/eco-go/database"
	"github.com/zainabmohammed9949/eco-go/middleware"
	routes "github.com/zainabmohammed9949/eco-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/Cartcheckout", app.BuyFromCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
