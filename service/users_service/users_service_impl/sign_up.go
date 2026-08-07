package users_service_imp

import (
	"github.com/Zeeshanahmed-builds/ChatApp/models"
	"errors"
	"fmt"
)

func (a *AuthUsers_Imp) SignUp(users *models.User) error {

	err := a.users.SignUp(users)

	if err != nil {
		fmt.Println("Error signing up user:", err)
		return errors.New("Internal Server Error")
	}

	return nil
}
