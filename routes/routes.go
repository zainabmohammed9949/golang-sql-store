package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zainabmohammed9949/eco-go/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Post("/users/signup", controllers.SignUp())
	incomingRoutes.Post("/users/login", controllers.Login())
	incomingRoutes.Post("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
