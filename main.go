package main

import (
	"restapi/handlers"
	"restapi/middleware"

	"restapi/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabase()
	r := gin.Default()
	r.GET("/debug-token", handlers.DebugTokenHandler)

	r.GET("/check-token", handlers.CheckTokenHandler)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.GET("/me", middleware.JWTAuthMiddleware(), handlers.MeHandler)
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.POST("/refresh", handlers.Refresh)
	r.POST("/chats", middleware.JWTAuthMiddleware(), handlers.CreateChat)
	r.GET("/chats", middleware.JWTAuthMiddleware(), handlers.GetChats)
	r.POST("/chats/:id/messages", middleware.JWTAuthMiddleware(), handlers.CreateMessage)

	r.Run(":8080")
}
