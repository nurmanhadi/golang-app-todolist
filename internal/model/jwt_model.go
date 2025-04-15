package model

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaim struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
