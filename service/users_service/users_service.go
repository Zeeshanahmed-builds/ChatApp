package users_service

import (
	"github.com/Zeeshanahmed-builds/ChatApp/models"
)

type AuthUsers interface {
	Login(login *models.Login) (string, error)
	SignUp(users *models.User) error
}
