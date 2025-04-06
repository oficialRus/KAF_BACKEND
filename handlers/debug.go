package handlers

import (
	"net/http"
	"restapi/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func DebugTokenHandler(c *gin.Context) {
	token, err := utils.GenerateToken(1, "access", 15*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Не удалось создать токен",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
