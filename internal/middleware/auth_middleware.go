package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Use .env to your custom 'secret_key'
//var jwtKey = []byte("your_secret_key")
var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// AuthMiddleware validates JWT tokens from the Authorization header
// and extracts the user ID into the request context.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header from the request
		authHeader := c.GetHeader("Authorization")

		// If no Authorization header is provided, reject the request
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		// Remove "Bearer " prefix from the token string
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Prepare a map to store JWT claims (payload data)
		// Parse and validate the token using the provided key
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		// If parsing fails or token is invalid, reject the request
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Store userID in Gin context for use in downstream handlers
		userID := int(claims["user_id"].(float64))
		c.Set("userID", userID)

		// Continue to the next middleware/handler
		c.Next()
	}
}