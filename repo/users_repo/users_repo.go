package users_repo

import "chat-app/models"

type UsersRepo interface {
	SignUp(users *models.Users) error
	Login(login *models.Login) (*models.Users, error)
}
