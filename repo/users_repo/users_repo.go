package users_repo

import "chat-app/models"

type UsersRepo interface {
	SignUp(user *models.User) error
	Login(email string) (*models.User, error)
}
