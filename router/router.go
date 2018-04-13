package router

import (
	"path"
	"path/filepath"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// Init - router collection and controller dispatch
func Init() {
	// default config include the recover and logger middleware
	router := gin.Default()
	// swaggerGroup(router)
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	v1Prefix := "/api/v1"
	todoGroup(router, v1Prefix)
	authGroup(router, v1Prefix)
	// SPA fallback router
	router.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./public/index.html")
		} else {
			c.File("./public" + path.Join(dir, file))
		}

	})
	// router.Run()

	router.RunTLS(":8080", "./cert/server.pem", "./cert/server.key")
}
