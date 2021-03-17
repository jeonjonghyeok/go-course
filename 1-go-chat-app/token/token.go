package token

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const mySigningKey = "32iazLZ3hD4aH4EKjRkEo3is"

type customClaims struct {
	UserID int `json:"uid"`
	jwt.StandardClaims
}

func New(userid int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims{
		UserID: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			Issuer:    "Learningspoons Chat",
		},
	})

	return token.SignedString([]byte(mySigningKey))
}

func Parse(token string) (userid int, err error) {
	parsed, err := jwt.ParseWithClaims(token, &customClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(mySigningKey), nil
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
