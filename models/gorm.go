package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

type Message struct {
	gorm.Model

	SenderID   uint   `json:"sender_id" gorm:"not null"`
	ReceiverID uint   `json:"receiver_id" gorm:"not null"`
	Message    string `json:"message" gorm:"type:text;not null"`
	IsRead     bool   `json:"is_read" gorm:"default:false"`
}
