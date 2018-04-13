package constants

//WrappedError is the common struct with err info
type WrappedError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type errKey string

// errKeys name
var (
	EmptyErr             errKey = "empty"
	ParamsErr            errKey = "ParamsErr"
	UserNotExistErr      errKey = "UserNotExist"
	TokenGenFailedErr    errKey = "TokenGenFailed"
	UserCreatedFailedErr errKey = "UserCreatedFailed"
	TokenExpireErr       errKey = "TokenExpire"
)

// ErrorEnums contains all the error type will response
var ErrorEnums map[errKey]WrappedError

func init() {
	ErrorEnums = make(map[errKey]WrappedError)
	ErrorEnums[EmptyErr] = WrappedError{}
	ErrorEnums[ParamsErr] = WrappedError{
		-100, "incorrect params",
	}
	ErrorEnums[UserNotExistErr] = WrappedError{
		-101, "User not exist",
	}
	ErrorEnums[TokenGenFailedErr] = WrappedError{
		-102, "Token generated failed",
	}
	ErrorEnums[UserCreatedFailedErr] = WrappedError{
		-103, "User record created failed",
	}
	ErrorEnums[TokenExpireErr] = WrappedError{
		-104, "Token expired, please refresh",
	}
}
