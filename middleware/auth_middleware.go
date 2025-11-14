package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Read Authorization header

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization Header missing"})

			c.Abort()
			return
		}
		// Expected Bearer Token

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token Format"})
			c.Abort()
			return
		}

		// Parse token

		secret := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Inavalid or expired token"})

			c.Abort()
			return
		}

		// Extract claims

		claims := token.Claims.(jwt.MapClaims)

		userID := uint(claims["user_id"].(float64))

		// save user ID for use in controllers

		c.Set("user_id", userID)

		c.Next()
	}
}
