package handler

import (
	"chat-app/service/users_service"
	"chat-app/service/message_service"
)

type AuthHandler struct {
	AuthUsers 		users_service.AuthUsers
	MessageService 	message_service.MessageService
}

func NewAuthHandler (
	users 	 users_service.AuthUsers,
	messageService message_service.MessageService,
	) *AuthHandler{
	return &AuthHandler{
		AuthUsers: users,
		MessageService: messageService,
	}
}