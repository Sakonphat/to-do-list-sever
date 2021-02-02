package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sever/services"
)

func Authentication() gin.HandlerFunc {

	return func(c *gin.Context) {

		accessToken, err := services.ParseJwtToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : http.StatusUnauthorized,
				"success" : false,
				"message" : "unauthorized",
			})
			return
		}

		user, userErr := services.GetUser(accessToken)
		if userErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : http.StatusUnauthorized,
				"success" : false,
				"message" : "unauthorized",
			})
			return
		}

		c.Set("user", user)

		c.Next()
	}

}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func ErrorHandler(c *gin.Context) {
	if len(c.Errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code" : http.StatusInternalServerError,
			"success" : false,
			"message" : c.Errors,
		})
		return
	}

	c.Next()
}