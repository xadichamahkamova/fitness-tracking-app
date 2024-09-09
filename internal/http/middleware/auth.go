package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	t "github.com/xadichamahkamova/fitness-tracking-app/internal/http/token"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token := ctx.GetHeader("Authorization")
		log.Println("Authorization header:", token)
		url := ctx.Request.URL.Path
		log.Println("Request URL:", url)

		if strings.Contains(url, "swagger") || url == "/users/register" || url == "/users/login" {
			ctx.Next()
			return
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Authorization header is missing",
			})
			return
		}

		// Bearer prefix borligini tekshirish
		if !strings.HasPrefix(token, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Authorization token is missing Bearer prefix",
			})
			return
		}

		// Bearer prefix ni ochirish
		token = strings.TrimPrefix(token, "Bearer ")

		userID, err := t.ExtractClaim(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		ctx.Set("userID", userID)
		ctx.Next()
	}
}
