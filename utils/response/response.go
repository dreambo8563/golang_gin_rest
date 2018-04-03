package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"vincent.com/golangginrest/constants"
)

// Response is a wrapper with extra info
func Response(c *gin.Context, data interface{}, e constants.WrappedError, code int) {
	if code == 0 {
		code = http.StatusOK
	}

	c.JSON(code, gin.H{
		"data":  data,
		"error": e,
	})
}
