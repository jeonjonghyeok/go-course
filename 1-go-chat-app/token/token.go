package token

import (
	"github.com/dgrijalva/jwt-go"
)
func newToken() {
	jwt.NewWithClaims(jwt.SigningMethodHS256).
}
