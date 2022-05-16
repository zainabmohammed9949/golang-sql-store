package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zainabmohammed9949/eco-go/models"
	"www.github.com/gin-gonic/gin"
)

func HashPassword(password string) string {

}

func VerfyPassward(userPassword string, givenPassword string) (bool, string) {

}

func Signup() gin.handlerfunc {

	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON{http.StatusBadRequest, gin.H{"error": err.Error()}}
			return
		}
		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}
		count, err := UserCollection, CountDocument(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON{http.StatusInternalServerError, gin.H{"error": err}}
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		}
		count, err = UserCollection.CountDocument(ctx, bson.H{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON{http.StatusInternalServerError, gin.H{"error": err}}
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "this phone is already in use"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password
		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		token, refreshtoken := generate.TokenGenerator(*user.Email, user.First_Name, user.Last_Name, user.User_ID)
		user.Token = &token
		user.Refresh_Token = &refreshtoken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)
		_, inserterr := UserCollection.InsertOne(ctx, user)
		if inserterr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "the user did not get created"})
			return
		}
		defer cancel()

		c.JSON(http.StatusCreated, "Successfully signed in!")
	}
}

func Login() gin.handlerfunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON{http.StatusBadRequest, gin.H{"error": err}}
			return
		}
		err := UserCollection.FinedOne(ctx, bson.H{"email": user.Email}).Decode(&founduser)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "The Email or Password is not correct"})
			return
		}
		PasswordIsValid, msg = VerfyPassward(*user.Password, *founduser.Password)
		if PasswordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": emsg})
			fmt.Println(msg)
			return
		}
		token, refreshtoken, _ := generate.TokenGenerator(*founduser.Email, founduser.First_Name, founduser.Last_Name, founduser.User_ID)
		defer cancel()
		generate.UpdatedAllTokens(token, refreshtoken, founduser.User_ID)
		c.JSON(http.StatusFound, founduser)
	}

}

func searchProduct() gin.handlerfunc {

}

func searchProductByQuery() gin.HandlerFunc {

}
