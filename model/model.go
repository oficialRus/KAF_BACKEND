package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Pass  int
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
