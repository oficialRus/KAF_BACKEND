package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    string `json:"id "`
	Email string `gorm:"unique"`
	Pass  int    `json:"password"`
	Chats []Chat `gorm:"foreignKey:UserID"`
}

type Chat struct {
	gorm.Model
	UserID   int
	Messages []Message `gorm:"foreignKey:ChatID"`
}

type Message struct {
	gorm.Model
	ChatID    int
	Content   string
	Role      string
	SeachType string
}
