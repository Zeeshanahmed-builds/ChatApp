package handler

import (
	"chat-app/service/users_service"
)

type AuthHandler struct {
	AuthUsers 	users_service.AuthUsers
}

func NewAuthHandler (
	users 	 users_service.AuthUsers,
	
	) *AuthHandler{
	return &AuthHandler{
		AuthUsers: users,
	}
}