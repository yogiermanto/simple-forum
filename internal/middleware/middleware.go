package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-forum/internal/configs"
	"simple-forum/internal/pkg/jwt"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		headers := strings.Split(header, " ")
		if len(headers) != 2 {
			c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		if headers[0] != "Bearer" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		}

		if headers[1] == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userID, username, err := jwt.ValidateToken(headers[1], secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("userID", userID)
		c.Set("username", username)
		c.Next()
	}
}

func AuthRefreshTokenMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		headers := strings.Split(header, " ")
		if len(headers) != 2 {
			c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		if headers[0] != "Bearer" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		}

		if headers[1] == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userID, username, err := jwt.ValidateTokenWithoutExpires(headers[1], secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("userID", userID)
		c.Set("username", username)
		c.Next()
	}
}
