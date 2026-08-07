package handler

import (
	"github.com/Zeeshanahmed-builds/ChatApp/service/message_service"
	"github.com/Zeeshanahmed-builds/ChatApp/service/users_service"
)

type AuthHandler struct {
	AuthUsers      users_service.AuthUsers
	MessageService message_service.MessageService
}

func NewAuthHandler(
	users users_service.AuthUsers,
	messageService message_service.MessageService,
) *AuthHandler {
	return &AuthHandler{
		AuthUsers:      users,
		MessageService: messageService,
	}
}
