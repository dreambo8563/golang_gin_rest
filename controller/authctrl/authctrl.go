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
		Token string `json:"token"`
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
	jwt.SaveToken(c, token)
	log.Infoln("token", token)
	response.Response(c, tokenResp{token}, constants.ErrorEnums["empty"], http.StatusOK)
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
	jwt.SaveToken(c, token)
	response.Response(c, tokenResp{token}, constants.ErrorEnums["empty"], http.StatusOK)
}

// Refresh will parse the expired token and return a fresh one
func Refresh(c *gin.Context) {
	token, err := jwt.GetToken(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	claims, exp, err := jwt.Parse(token)
	if exp {
		userID := claims["userID"].(string)
		token, err = jwt.New(userID)
		if err != nil {
			response.Response(c, nil, constants.ErrorEnums["TokenGenFailed"], 0)
			return
		}
		jwt.SaveToken(c, token)
		response.Response(c, tokenResp{token}, constants.ErrorEnums["empty"], http.StatusOK)
		return
	}
	c.AbortWithStatus(http.StatusUnauthorized)
}

// Test is the handle to test jwt middleware
func Test(c *gin.Context) {
	userID, exist := c.Get("userID")
	if exist {
		log.Info("userID: ", userID)
	} else {
		log.Panic("AuthMiddleware fail")
		return
	}
	response.Response(c, nil, constants.ErrorEnums["empty"], http.StatusOK)
}
