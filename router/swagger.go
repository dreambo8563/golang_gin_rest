package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// make different group for each module
func swaggerGroup(r *gin.Engine) {
	swaggerRouter := r.Group("/swagger")
	{
		swaggerRouter.GET("/", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

}
