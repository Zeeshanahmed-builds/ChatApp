package utils

import (
	"github.com/golang-jwt/jwt"
)

var key = []byte("mysecretkey12345")

func GenerateToken(email string, user_id int) (string, error) {

	// 16 bytes key for AES-128

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":   email,
			"user_id": user_id,
		},
	)

	// claims := token.Claims.(jwt.MapClaims)
	// User_id := int(claims["user_id"].(float64))

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, err
}