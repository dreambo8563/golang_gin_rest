package router

import (
	"github.com/gin-gonic/gin"
	"vincent.com/golangginrest/controller"
)

// InitRouter - router collection and controller dispatch
func InitRouter() {
	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", controller.CreateTodo)
		v1.GET("/", controller.FetchAllTodo)
		v1.GET("/:id", controller.FetchSingleTodo)
		v1.PUT("/:id", controller.UpdateTodo)
		v1.DELETE("/:id", controller.DeleteTodo)
	}
	router.Run()
}
