package users_service_imp

import (
	"chat-app/repo/users_repo"
	"chat-app/service/users_service"
	
)

type AuthUsers_Imp struct{
	users		users_repo.UsersRepo
}


func NewAuthUsers(input NewAuthUsersImp)users_service.AuthUsers{
	return &AuthUsers_Imp{
		users: input.Users,
	}
}

type NewAuthUsersImp struct{
	Users  users_repo.UsersRepo
}