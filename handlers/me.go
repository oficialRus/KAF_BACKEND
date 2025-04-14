package handlers

import (
	"net/http"
	"restapi/database"
	"restapi/model"

	"github.com/gin-gonic/gin"
)

func MeHandler(controller *gin.Context) {
	userID, exists := controller.Get("user_id")
	if !exists {
		controller.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось получить user_id"})
		return
	}

	uid, ok := userID.(uint)
	if !ok {
		controller.JSON(http.StatusInternalServerError, gin.H{"error": "user_id имеет неверный тип"})
		return
	}

	var user model.User
	if err := database.DB.First(&user, uid).Error; err != nil {
		controller.JSON(http.StatusInternalServerError, gin.H{"error": "пользователь не найден"})
		return
	}

	controller.JSON(http.StatusOK, gin.H{
		"message": "Вы авторизованы",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
