package routes

import (
	"github.com/gin-gonic/gin"
	"sever/controllers"
	"sever/middlewares"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(middlewares.Cors())

	api := "/api"

	v1 := r.Group(api+"/v1")
	{
		v1.POST("/register", controllers.Register)
		v1.POST("/login", controllers.Login)
		v1.POST("/logout", controllers.Logout)

		v1.POST("/task", middlewares.Authentication(), controllers.CreateTask)
		v1.GET("/task/:uuid", middlewares.Authentication(), controllers.GetTask)
		v1.POST("/tasks", middlewares.Authentication(), controllers.GetAllTask)
		v1.PUT("/edit", middlewares.Authentication(), controllers.EditTask)
		v1.PUT("/complete/:uuid", middlewares.Authentication(), controllers.CompleteTask)
		v1.PUT("/undo/:uuid", middlewares.Authentication(), controllers.UndoTask)
		v1.DELETE("/delete/:uuid", middlewares.Authentication(), controllers.DeleteTask)
		v1.DELETE("/delete-all", middlewares.Authentication(), controllers.DeleteAllTask)
	}

	return r
}