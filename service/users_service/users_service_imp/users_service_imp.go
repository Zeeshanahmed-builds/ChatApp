package users_service_imp

import (
	"chat-app/repo"
	"chat-app/service/users_service"
	
)

type AuthUsers_Imp struct{
	users		repo.Users
}


func NewAuthUsers(input NewAuthUsersImp)users_service.AuthUsers{
	return &AuthUsers_Imp{
		users: input.Users,
	}
}

type NewAuthUsersImp struct{
	Users  repo.Users
}