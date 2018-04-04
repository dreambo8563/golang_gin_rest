package authctrl

import (
	"fmt"
	"net/http"

	"vincent.com/golangginrest/constants"
	"vincent.com/golangginrest/service/jwt"

	"github.com/gin-gonic/gin"
	"vincent.com/golangginrest/model"
	"vincent.com/golangginrest/utils/logger"
	"vincent.com/golangginrest/utils/response"
)

var log = logger.Log()

type (
	tokenResp struct {
		token string
	}
)

// Login - validate login info and return token
func Login(c *gin.Context) {
	var user model.UserModel
	err := c.ShouldBind(&user)

	// bind check
	if err != nil {
		log.Errorln(err.Error())
		response.Response(c, nil, constants.ErrorEnums["ParamsErr"], 0)
		return
	}
	// check user exist
	user, exist := model.FindByInfo(user.Phone, user.Password)
	if !exist {
		log.Infoln("user not exist")
		response.Response(c, nil, constants.ErrorEnums["UserNotExist"], 0)
		return
	}
	// generate token by userid
	token, err := jwt.New(fmt.Sprint(user.ID))
	// log.Infoln(token, err.Error())
	if err != nil {
		log.Infoln("TokenGenFailed")
		response.Response(c, nil, constants.ErrorEnums["TokenGenFailed"], 0)
		return
	}
	c.SetCookie("token", token, 3600, "/", "", true, true)
	log.Infoln("token", token)
	response.Response(c, gin.H{
		"token": token,
	}, constants.ErrorEnums["empty"], http.StatusOK)
}

// Register will create the user record
func Register(c *gin.Context) {
	var user model.UserModel
	err := c.ShouldBind(&user)
	// bind check
	if err != nil {
		log.Errorln(err.Error())
		response.Response(c, nil, constants.ErrorEnums["ParamsErr"], 0)
		return
	}
	err = user.New()
	if err != nil {
		response.Response(c, nil, constants.ErrorEnums["UserCreatedFailed"], 0)
		return
	}
	// generate token by userid
	token, err := jwt.New(fmt.Sprint(user.ID))
	// log.Infoln(token, err.Error())
	if err != nil {
		response.Response(c, nil, constants.ErrorEnums["TokenGenFailed"], 0)
		return
	}

	c.SetCookie("token", token, 3600, "/", "", true, true)

	response.Response(c, gin.H{
		"token": token,
	}, constants.ErrorEnums["empty"], http.StatusOK)
}
