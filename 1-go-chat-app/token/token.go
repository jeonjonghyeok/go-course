package token

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type customClaims struct {
	UserID int `json:"uid`
	jwt.StandardClaims
}

func New(userid int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims{
		UserID: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			Issuer:    "jjh",
		},
	})

	return token.SignedString(([]byte("jjh")))
}

func Parse(token string) (userid int, err error) {
	parsed, err := jwt.ParseWithClaims(token, customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("jjh"), nil
	})
	if err != nil {
		return 0, err
	}

	if !parsed.Valid {
		return 0, errors.New("token is invalid")
	}

	if c, ok := parsed.Claims.(*customClaims); ok {
		return c.UserID, nil
	}
	return 0, errors.New("token is invalid")
}
