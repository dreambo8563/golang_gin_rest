package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"vincent.com/golangginrest/constants"
)

type resp struct {
	Data    interface{}            `json:"data"`
	Error   constants.WrappedError `json:"error"`
	Success bool                   `json:"success"`
}

// Response is a wrapper with extra info
func Response(c *gin.Context, data interface{}, e constants.WrappedError, code int) {
	if code == 0 {
		code = http.StatusOK
	}
	success := true

	if code >= 400 {
		success = false
	}

	c.JSON(code, resp{
		data,
		e,
		success,
	})
}
