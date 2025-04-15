package service

import (
	"golang-app-todolist/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func JwtGenerateAccesToken(username string, key []byte) (string, error) {
	claims := model.JwtCustomClaim{
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
func JwtVerifyToken(tokenString string, key []byte) (*model.JwtCustomClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(*model.JwtCustomClaim)
	claimType := &model.JwtCustomClaim{
		Username: claims.Username,
	}
	return claimType, nil
}
