package constants

//WrappedError is the common struct with err info
type WrappedError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// ErrorEnums contains all the error type will response
var ErrorEnums = map[string]WrappedError{
	"empty": WrappedError{},
	"ParamsErr": WrappedError{
		-100, "incorrect params",
	},
	"UserNotExist": WrappedError{
		-101, "User not exist",
	},
	"TokenGenFailed": WrappedError{
		-102, "Token generated failed",
	},
	"UserCreatedFailed": WrappedError{
		-103, "User record created failed",
	},
	"TokenExpire": WrappedError{
		-104, "Token expired, please refresh",
	},
}
