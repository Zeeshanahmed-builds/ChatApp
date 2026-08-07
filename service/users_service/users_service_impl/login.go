package users_service_imp

import (
	"errors"
	"fmt"
	"github.com/Zeeshanahmed-builds/ChatApp/models"
	"github.com/Zeeshanahmed-builds/ChatApp/utils"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (a *AuthUsers_Imp) Login(login *models.Login) (string, error) {

	user, err := a.users.Login(login.Email)
	if err != nil {
		fmt.Println("Error logging in user:", err)
		return "", errors.New("Invalid email or password")
	}

	userPass := []byte(user.Password)
	loginPass := []byte(login.Password)
	passErr := bcrypt.CompareHashAndPassword(userPass, loginPass)
	if passErr != nil {
		fmt.Println("Error in password", passErr)
		log.Println(passErr)
		return "", passErr
	}

	token, err := utils.GenerateToken(user.Email, int(user.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}
