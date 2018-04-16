package router

import (
	"path"
	"path/filepath"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Init - router collection and controller dispatch
func Init() {
	env := viper.Get("GO_ENV")
	// default config include the recover and logger middleware
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	// swaggerGroup(router)
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Static("/.well-known/acme-challenge", "./.well-known/acme-challenge")
	v1Prefix := "/api/v1"
	todoGroup(router, v1Prefix)
	authGroup(router, v1Prefix)
	// SPA fallback router
	router.NoRoute(func(c *gin.Context) {
		_, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./public/index.html")
		} else {
			router.Static("/", "./public")
			// c.File("./public" + path.Join(dir, strings.Split(file, "?")[0]))
		}

	})
	// router.Run(":80")
	// log.Fatal(autotls.Run(router, "dreambo8563.tech"))

	if env == "production" {
		router.RunTLS(":443", "./cert/fullchain.pem", "./cert/privkey.pem")
	} else {
		router.RunTLS(":8080", "./cert/server.pem", "./cert/server.key")
	}

}
