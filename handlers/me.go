package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MeHandler(controller *gin.Context) {
	userID, exists := controller.Get("user_id")
	if !exists {
		controller.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось получить user_id"})
		return
	}
	controller.JSON(http.StatusOK, gin.H{
		"message": "Вы авторизованы",
		"user_id": userID,
	})
}
