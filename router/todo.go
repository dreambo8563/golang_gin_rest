package router

import (
	"github.com/gin-gonic/gin"
	"vincent.com/golangginrest/controller/todoctrl"
)

// make different group for each module
func todoGroup(r *gin.Engine, prefix string) {

	todoRouter := r.Group(prefix + "/todo")
	{
		todoRouter.POST("/", todoctrl.CreateTodo)
		todoRouter.GET("/", todoctrl.FetchAllTodo)
		todoRouter.GET("/:id", todoctrl.FetchSingleTodo)
		todoRouter.PUT("/:id", todoctrl.UpdateTodo)
		todoRouter.DELETE("/:id", todoctrl.DeleteTodo)
	}

}
