package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name  string `json:"name"`
	Email string `gorm:"uniqueIndex" json:"email"`
	Pass  string `json:"password"`
	Chats []Chat `gorm:"foreignKey:UserID"`
}

type Chat struct {
	gorm.Model
	UserID   uint
	Title    string
	Messages []Message `gorm:"foreignKey:ChatID"`
}

type Message struct {
	gorm.Model
	ChatID     int
	Content    string
	Role       string
	SearchType string
}
