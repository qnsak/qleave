package domain

import jwt "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	ID       string
	Username string
	jwt.StandardClaims
}
