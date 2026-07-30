package users_service_imp
import (
	"fmt"
	"errors"
	"chat-app/models"
	
)

func (a *AuthUsers_Imp) SignUp(users *models.User) error {
	
	 err:=a.users.SignUp(users)

	if err != nil {
		fmt.Println("Error signing up user:", err)
		return errors.New("Internal Server Error")
	}
		

	return nil
}

