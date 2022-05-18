package controller	"go.mongodb.org/mongo-driver/bson/primitive"
s

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
	"errors"
	"log"
	"errors"
	"context"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"

)
    
    type Application struct{
		prodCollection * mongo.Collection
		UserCollection * mongo.Collection
	}
	func NewApplication(prodCollection, userCollection * mongo.Collection )*Application{
		return &Application{
			prodCollection: prodCollection,
			UserCollection: userCollection,
		}
	}

func(app *Application) AddToCart() gin.HandlerFunc {
return func(c *gin.Context) {
	productQueryID := c.Query("id")
	if productQueryID ==""{ 
		log.Println("product id is empty")
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("product is empty"))
		return
	}
    userQueryID := c.Query("userId")
	if userQueryID ==""{ 
		log.Println("user id is empty")
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("user  id is empty"))
		return
	}

	productID, err := primitive.ObjectIDFormatHex(productQueryID)
	if err!= nil{
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
     
    var ctx, cancel = context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	err = database.AddProductToCart(ctx, app.prodCollection, app.UserCollection, productID, userQueryID)
    if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(200, "Successfully added to the cart")
	}
}
func RemoveItem() gin.HandlerFunc {

}
func GetItemFromCart() gin.HandlerFunc {

}
func BuyFromCart() gin.HandlerFunc {

}
func InstantBuy() gin.HandlerFunc {

}
