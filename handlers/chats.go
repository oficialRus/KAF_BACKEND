package handlers

import (
	"net/http"
	"restapi/database"
	"restapi/model"

	"github.com/gin-gonic/gin"
)

type CreateChatRequest struct {
	Title string `json:"title"`
}

// ///POST
func CreateChat(controller *gin.Context) {
	userID, exists := controller.Get("user_id")
	if !exists {
		controller.JSON(http.StatusUnauthorized, gin.H{"error": "user_id не найден"})
		return
	}

	uid, ok := userID.(uint)
	if !ok {
		controller.JSON(http.StatusInternalServerError, gin.H{"error": "user_id неверного типа"})
		return
	}

	var req struct {
		Title string `json:"title"`
	}

	if err := controller.ShouldBindJSON(&req); err != nil {
		controller.JSON(http.StatusBadRequest, gin.H{"error": "неверный формат запроса"})
		return
	}

	chat := model.Chat{
		UserID: uid,
		Title:  req.Title,
	}

	if err := database.DB.Create(&chat).Error; err != nil {
		controller.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось создать чат"})
		return
	}

	controller.JSON(http.StatusCreated, gin.H{"chat_id": chat.ID})
}

/////GET

func GetChats(controller *gin.Context) {
	userID, exists := controller.Get("user_id")
	if !exists {
		controller.JSON(http.StatusUnauthorized, gin.H{"error": "user_id не найден"})
		return

	}
	uid, ok := userID.(uint)
	if !ok {
		controller.JSON(http.StatusInternalServerError, gin.H{"error": "user_id неверного типа"})
		return
	}
	var chats []model.Chat
	if err := database.DB.Where("user_id=?", uid).Find(&chats).Error; err != nil {
		controller.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось получить чаты"})
		return
	}
	controller.JSON(http.StatusOK, chats)

}
