package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userId string, username string, jwtSecret string, expiresAt time.Duration) (string, error) {
	claims := Claims{
		Id:       userId,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "jejakmakan",
			Subject:   userId,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresAt)),
			ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
}

func ParseToken(tokenString string, jwtSecret string) (string, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	}

	parsedToken, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return "", err
	}
	if !parsedToken.Valid {
		return "", errors.New("invalid token")
	}
	return parsedToken.Raw, nil
}
