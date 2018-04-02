package router

import (
	"path"

	"github.com/gin-gonic/gin"
	// import the generated doc package
	_ "vincent.com/golangginrest/docs"
)

// Init - router collection and controller dispatch
func Init() {
	// default config include the recover and logger middleware
	router := gin.Default()
	v1Prefix := "/api/v1"
	swaggerGroup(router)
	todoGroup(router, v1Prefix)
	// SPA fallback router
	router.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		if file == "" {
			c.File("./public/index.html")
		} else {
			c.File("./public" + path.Join(dir, file))
		}

	})
	// router.Run()

	router.RunTLS(":8080", "./cert/server.pem", "./cert/server.key")
}
