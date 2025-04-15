package service

import (
	"golang-app-todolist/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func JwtGenerateAccesToken(username string, key []byte) (string, error) {
	claims := model.JwtCustomeClaim{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add((time.Hour * 24) * 7)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return ss, nil
}
