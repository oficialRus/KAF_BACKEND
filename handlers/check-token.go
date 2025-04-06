package handlers

import (
	"restapi/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckTokenHandler(controller *gin.Context) {
	authHeader := controller.GetHeader("Authorization")
	if authHeader == "" {
		controller.JSON(401, gin.H{"error": "Токен не передан"})
		return
	}
	tokenStr := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	claims, err := utils.ParseToken(tokenStr)
	if err != nil {
		controller.JSON(401, gin.H{"error": "Невалидный токен"})
		return
	}
	controller.JSON(200, gin.H{
		"user_id": claims.UserID,
		"type":    claims.Type,
	})
}
