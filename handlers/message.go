package handlers

import (
	"net/http"
	"restapi/database"
	"restapi/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MessageRequest struct {
	Content    string `json:"content" binding:"required"`
	Role       string `json:"role" binding:"required"`
	SearchType string `json:"search_type" binding:"required"`
}

func CreateMessage(c *gin.Context) {
	chatIDParam := c.Param("id")

	chatIDUint, err := strconv.ParseUint(chatIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный chat_id"})
		return
	}

	chatID := uint(chatIDUint)

	userIDRaw, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	userID, ok := userIDRaw.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка типа user_id"})
		return
	}

	var chat model.Chat
	if err := database.DB.Where("id=? AND user_id=?", chatID, userID).First(&chat).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Чат не найден или не принадлежит вам"})
		return
	}

	// Всё хорошо — можно продолжать дальше
	c.JSON(http.StatusOK, gin.H{
		"chat_id": chat.ID,
	})

	var req MessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный формат данных"})
		return
	}
	message := model.Message{
		ChatID:     chat.ID,
		Content:    req.Content,
		Role:       req.Role,
		SearchType: req.SearchType,
	}
	if err := database.DB.Create(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось сохранить сообщение"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message_id": message.ID})
}
