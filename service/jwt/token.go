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
	UserID string `json:"userId"`
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
	tokenExp := time.Now().Add(time.Hour * 1).Unix()
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
func Parse(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secretString), nil
	})
	// actually we will get err only when the token expire
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		return claims, nil
	}
	return nil, fmt.Errorf("the token is invalid")
}
