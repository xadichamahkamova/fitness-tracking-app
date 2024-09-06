package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	t "github.com/xadichamahkamova/fitness-tracking-app/internal/http/token"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		userID, err := t.ExtractClaim(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}

