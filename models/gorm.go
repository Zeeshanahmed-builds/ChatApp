package models

import (
	"gorm.io/gorm"
)

// User table
type User struct {	
	gorm.Model // ID, CreatedAt, UpdatedAt, DeletedAt

	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

// // Login request (not a database table)
// type Login struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// Message table
type Message struct {
	gorm.Model // ID, CreatedAt, UpdatedAt, DeletedAt

	SenderID   uint   `json:"sender_id" gorm:"not null"`
	ReceiverID uint   `json:"receiver_id" gorm:"not null"`
	Message    string `json:"message" gorm:"type:text;not null"`
	IsRead     bool   `json:"is_read" gorm:"default:false"`
}