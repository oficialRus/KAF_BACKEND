package handlers

import (
	"net/http"
	"restapi/database"
	"restapi/model"
	"restapi/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Register(controller *gin.Context) {
	var req RegisterRequest

	if err := controller.ShouldBindJSON(&req); err != nil {
		controller.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return

	}
	var existing model.User
	if err := database.DB.Where("email=?", req.Email).First(&existing).Error; err == nil {
		controller.JSON(http.StatusConflict, gin.H{"error": "Email уже используется"})
		return

	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		controller.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка хеширования пароля"})
		return
	}
	user := model.User{
		Name:  req.Name,
		Email: req.Email,
		Pass:  hashedPassword,
	}
	if err := database.DB.Create(&user).Error; err != nil {
		controller.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении пользователя"})
		return
	}
	controller.JSON(http.StatusCreated, gin.H{"message": "Пользователь успешно зарегистрирован"})
}

func Login(controller *gin.Context) {
	var req LoginRequest
	if err := controller.ShouldBindJSON(&req); err != nil {
		controller.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	var user model.User
	if err := database.DB.Where("email=?", req.Email).First(&user).Error; err != nil {
		controller.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
		return
	}
	accessToken, err := utils.GenerateToken(user.ID, "access", 15*time.Minute)
	if err != nil {

		controller.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания access токена"})
		return
	}
	refreshToken, err := utils.GenerateToken(user.ID, "refresh", 7*24*time.Hour)
	if err != nil {
		controller.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания refresh токена"})
		return
	}
	controller.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func Refresh(c *gin.Context) {
	var req RefreshRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат"})
		return
	}

	claims, err := utils.ParseToken(req.RefreshToken)
	if err != nil || claims.Type != "refresh" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный или просроченный refresh токен"})
		return
	}

	accessToken, err := utils.GenerateToken(claims.UserID, "access", 15*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать токен"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}
