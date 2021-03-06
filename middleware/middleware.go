package middleware

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	generate "github.com/zainabmohammed9949/golang-sql-store/tokens"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		userToken := c.Request.Header.Get("token")
		if userToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No authentication header provided")})
			c.Abort()
			return
		}
		claims, err := generate.ValidateToken(userToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("user_id", claims.Uid)
		c.Next()
	}

}
