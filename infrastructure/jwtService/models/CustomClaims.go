package models

import "github.com/golang-jwt/jwt/v4"

type CustomClaims struct {
	jwt.RegisteredClaims
	Role  string `json:"role"`
	Scope string `json:"scope"`
}
