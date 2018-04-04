package middleware

import (
	"net/http"

	"vincent.com/golangginrest/utils/logger"

	"vincent.com/golangginrest/service/jwt"

	"github.com/gin-gonic/gin"
)

var log = logger.Log()

//AuthMiddleware will deal with the jwt token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, exp, err := jwt.Parse(token)
		if exp {
			c.AbortWithStatus(http.StatusNonAuthoritativeInfo)
			return
		}
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// log.Infoln(claims["userID"])
		c.Set("userID", claims["userID"])
		c.Next()
	}
}
