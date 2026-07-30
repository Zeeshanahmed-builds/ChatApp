package users_service

import (
	"chat-app/models"
)

type AuthUsers interface {
	Login(login *models.Login) (string, error)
	SignUp(users *models.User) error
}
