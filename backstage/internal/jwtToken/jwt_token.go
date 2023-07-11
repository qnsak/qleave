package jwtToken

import (
	"fmt"
	"leaveManager/domain"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// 產生 jwt token
func GenerateToken(user domain.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3000 * time.Second)
	issuer := user.Name
	id := user.ID
	claims := domain.JwtClaims{
		ID:       id,
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("golang"))

	return token, err
}

// 更新 jwt token
func CreateRefreshToken() (refreshToken string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(600 * time.Second)
	issuer := "frank"
	claims := domain.JwtClaims{
		ID:       "10001",
		Username: "frank",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	rt, err := token.SignedString([]byte("golang"))
	if err != nil {
		return "", err
	}
	return rt, err
}

// 取得 user id
func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", err
	}

	id := claims["ID"].(string)

	return id, nil
}

// token 驗證
func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
