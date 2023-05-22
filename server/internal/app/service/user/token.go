package user

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(inputUser User) (string, error) {
	expiresAt := time.Now().Add(time.Minute * 30).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"account": inputUser.Account,
		"expire":  expiresAt,
		"status":  Hadlogined,
	})
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}
