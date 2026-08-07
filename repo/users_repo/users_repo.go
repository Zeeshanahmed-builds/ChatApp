package users_repo

import "github.com/Zeeshanahmed-builds/ChatApp/models"

type UsersRepo interface {
	SignUp(user *models.User) error
	Login(email string) (*models.User, error)
}
