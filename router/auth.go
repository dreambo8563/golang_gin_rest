package router

import (
	"github.com/gin-gonic/gin"
	"vincent.com/golangginrest/controller/authctrl"
	"vincent.com/golangginrest/middleware"
)

func authGroup(r *gin.Engine, prefix string) {
	authRouter := r.Group(prefix + "/auth")
	{
		authRouter.POST("/", authctrl.Login)
		authRouter.POST("/reg", authctrl.Register)
		authRouter.GET("/test", middleware.AuthMiddleware(), authctrl.Test)
	}
}
