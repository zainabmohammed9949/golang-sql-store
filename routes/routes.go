package routes

import (
	//"log"
	//"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zainabmohammed9949/golang-sql-store/controllers"
)

func SetUp(app *fiber.App) {

	app.Post("/api/products", controllers.AddProduct)
	app.Get("/api/products", controllers.GetProducts)
	app.Get("/api/products/:id", controllers.GetProductByID)
	//app.Put("/api/products", UpdateProduct)
	app.Delete("/api/products/:id", controllers.DeleteProduct)

	app.Post("/api/orders", controllers.CreateOrder)
	app.Get("/api/orders", controllers.GetOrders)
	app.Get("/api/orders/:id", controllers.GetOrderByID)
	app.Delete("/api/orders/:id", controllers.CancelOrder)

	app.Post("/api/users/signup", controllers.Signup)
	app.Post("/api/users/login", controllers.Login)
	app.Post("/api/users/cookie", controllers.LoginWithCookie)

	app.Get("/api/users/cookie", controllers.GetUserByCookie)
	app.Get("/api/users", controllers.GetUsers)
	app.Get("/api/users/:id", controllers.GetUserByID)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	app.Post("/api/sellers/signup", controllers.SellerSignup)
	app.Post("/api/sellers/login", controllers.SellerLogin)
	app.Get("/api/sellers/cookie", controllers.GetSellerByCookie)
	app.Post("/api/sellers/cookie", controllers.LoginWithCookie)

	app.Get("/api/sellers", controllers.GetSellers)
	app.Get("/api/sellers/:id", controllers.GetSellerByID)
	app.Put("/api/sellers/:id", controllers.UpdateSeller)
	app.Delete("/api/sellers/:id", controllers.DeleteSeller)

	app.Get("/api/users/address", controllers.GetAddresses)
	app.Post("/api/users/address", controllers.AddAddress)
	app.Get("/api/users/address/:id", controllers.GetAddressByID)
	app.Put("/api/users/address/:id", controllers.UpdateAddress)
	app.Delete("/api/users/address/:id", controllers.DeleteAddress)

}
