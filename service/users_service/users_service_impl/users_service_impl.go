package users_service_imp

import (
	"github.com/Zeeshanahmed-builds/ChatApp/repo/users_repo"
	"github.com/Zeeshanahmed-builds/ChatApp/service/users_service"
)

type AuthUsers_Imp struct {
	users users_repo.UsersRepo
}

func NewAuthUsers(input NewAuthUsersImp) users_service.AuthUsers {
	return &AuthUsers_Imp{
		users: input.Users,
	}
}

type NewAuthUsersImp struct {
	Users users_repo.UsersRepo
}
