package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func AuthMiddleware(hmacSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		tokenType, tokenString := strings.Split(token, " ")[0], strings.Split(token, " ")[1]
		if tokenType != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		email, err := parseJWTToken(tokenString, []byte(hmacSecret))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		log.Printf("Authenticated user: %s", email)
		ctx := context.WithValue(c, "email", email)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func parseJWTToken(token string, hmacSecret []byte) (string, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})

	if err != nil {
		return "", err
	} else if claims, ok := t.Claims.(*Claims); ok {
		return claims.Email, nil
	} else {
		return "", fmt.Errorf("invalid token")
	}

}
