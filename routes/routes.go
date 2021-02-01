package routes

import (
	"github.com/gin-gonic/gin"
	"sever/controllers"
	"sever/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := "/api"
	v1 := r.Group(api+"/v1")
	{
		v1.POST("/register", controllers.Register)
		v1.POST("/login", controllers.Login)
		v1.POST("/logout", controllers.Logout)

		v1.POST("/task", middlewares.Authentication(), middlewares.ErrorHandler, controllers.CreateTask)
		v1.GET("/task/:uuid", middlewares.Authentication(), middlewares.ErrorHandler, controllers.GetTask)
		v1.GET("/tasks", middlewares.Authentication(), middlewares.ErrorHandler, controllers.GetAllTask)
		v1.PUT("/edit", middlewares.Authentication(), middlewares.ErrorHandler, controllers.EditTask)
		v1.PUT("/complete/{uuid}", middlewares.Authentication(), middlewares.ErrorHandler, controllers.CompleteTask)
		v1.PUT("/undo/{uuid}", middlewares.Authentication(), middlewares.ErrorHandler, controllers.UndoTask)
		v1.DELETE("/delete/{uuid}", middlewares.Authentication(), middlewares.ErrorHandler, controllers.DeleteTask)
		v1.DELETE("/delete-all", middlewares.Authentication(), middlewares.ErrorHandler, controllers.DeleteAllTask)
	}
	return r
}