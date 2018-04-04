package jwt

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	secretString []byte
	issuerString string
)

// UserClaims is the type contain userID and standardClaims
type UserClaims struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}

func init() {
	// default values
	secretString = []byte("secret")
	issuerString = "jwt"
}

// Config will set the basic info to generate token with jwt
func Config(secret string, issuer string) {
	if secret != "" {
		secretString = []byte(secret)
	}
	if issuer != "" {
		issuerString = issuer
	}
}

// New return the generated token with default alg and set/default config
func New(userID string) (string, error) {
	tokenExp := time.Now().Add(time.Second * 30).Unix()
	claims := UserClaims{
		userID,
		jwt.StandardClaims{
			Issuer:    issuerString,
			ExpiresAt: tokenExp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secretString)
	// fmt.Printf("%v %v", ss, err)
	// Parse(ss)
	return ss, err
}

// Parse - parse the tokenString into token struct
func Parse(tokenString string) (jwt.MapClaims, bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secretString), nil
	})
	// actually we will get err only when the token expire
	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			return claims, false, nil
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, false, fmt.Errorf("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				return claims, true, nil
			}
			// return nil, true, fmt.Errorf("the token parse error")
		}
	} else {
		return nil, false, fmt.Errorf("Couldn't handle this token")
	}

	return nil, false, fmt.Errorf("Couldn't handle this token")
}