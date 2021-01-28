package routes

import (
	"github.com/gin-gonic/gin"
	"sever/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.POST("/register", controllers.Register)
		v1.POST("/login", controllers.Login)
	}
	return r
}