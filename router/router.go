package router

import (
	"github.com/gin-gonic/gin"
)

// Init - router collection and controller dispatch
func Init() {
	router := gin.Default()
	v1Prefix := "/api/v1"
	todoGroup(router, v1Prefix)
	router.Run()
}
