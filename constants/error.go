package constants

//WrappedError is the common struct with err info
type WrappedError struct {
	code int
	msg  string
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
}
