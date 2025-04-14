package handlers

import (
	"net/http"
	"restapi/database"
	"restapi/model"
	"strconv"

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

func DeleteChat(controller *gin.Context) {
	chatIDParam := controller.Param("id")
	chatIDUint, err := strconv.ParseUint(chatIDParam, 10, 64)
	if err != nil {
		controller.JSON(http.StatusBadRequest, gin.H{"error": "неверный chat_id"})
		return
	}
	chatID := uint(chatIDUint)

	userIDRaw, exists := controller.Get("user_id")
	if !exists {
		controller.JSON(http.StatusUnauthorized, gin.H{"error": "не авторизован"})
		return
	}
	userID, ok := userIDRaw.(uint)
	if !ok {
		controller.JSON(http.StatusInternalServerError, gin.H{"error": "неверный тип user_id"})
		return
	}
	var chat model.Chat
	if err := database.DB.Where("id = ? AND user_id = ?", chatID, userID).First(&chat).Error; err != nil {
		controller.JSON(http.StatusNotFound, gin.H{"error": "Чат не найден или не принадлежит вам"})
		return
	}
	if err := database.DB.Where("chat_id = ?", chatID).Delete(&model.Message{}).Error; err != nil {
		controller.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось удалить сообщения чата"})
		return
	}
	if err := database.DB.Delete(&chat).Error; err != nil {
		controller.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось удалить чат"})
		return
	}

}
