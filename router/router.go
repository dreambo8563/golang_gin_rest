package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	// import the generated doc package
	_ "vincent.com/golangginrest/docs"
)

// Init - router collection and controller dispatch
func Init() {
	router := gin.Default()
	// router.Use(gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1Prefix := "/api/v1"
	todoGroup(router, v1Prefix)
	router.Run()
}
