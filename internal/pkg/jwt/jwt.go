package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func CreateToken(id int64, username, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"exp":      time.Now().Add(time.Minute * 10).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Bearer %s", tokenStr), nil
}

func ValidateToken(tokenStr, secretKey string) (int64, string, error) {
	key := []byte(secretKey)
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	return int64(claims["id"].(float64)), claims["username"].(string), nil

}

func ValidateTokenWithoutExpires(tokenStr, secretKey string) (int64, string, error) {
	key := []byte(secretKey)
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}, jwt.WithoutClaimsValidation())
	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	return int64(claims["id"].(float64)), claims["username"].(string), nil

}
