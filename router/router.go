package router

import (
	"github.com/gin-gonic/gin"
	"vincent.com/golangginrest/controller/todo"
)

// InitRouter - router collection and controller dispatch
func InitRouter() {
	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", todoctrl.CreateTodo)
		v1.GET("/", todoctrl.FetchAllTodo)
		v1.GET("/:id", todoctrl.FetchSingleTodo)
		v1.PUT("/:id", todoctrl.UpdateTodo)
		v1.DELETE("/:id", todoctrl.DeleteTodo)
	}
	router.Run()
}
