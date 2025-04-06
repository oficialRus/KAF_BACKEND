package middleware

import (
	"net/http"
	"restapi/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(controller *gin.Context) {
		authHeader := controller.GetHeader("Authorization")
		if authHeader == "" {
			controller.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "отсутствует токен"})
			return

		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ParseToken(tokenStr)
		if err != nil || claims.Type != "access" {
			controller.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "токен недействителен"})
			return
		}
		controller.Set("user_id", claims.UserID)
		controller.Next()
	}

}
