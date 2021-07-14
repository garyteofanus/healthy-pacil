package model

import "github.com/golang-jwt/jwt"

type CustomClaim struct {
	Role Role `json:"role"`
	jwt.StandardClaims
}
