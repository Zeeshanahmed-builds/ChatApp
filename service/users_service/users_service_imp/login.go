package users_service_imp

import (
	"chat-app/models"
	"chat-app/utils"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (a *AuthUsers_Imp) Login(login *models.Login) (string, error) {

	Resp, err := a.users.Login(login)
	if err != nil {
		fmt.Println("Error logging in user:", err)
		return "", errors.New("Invalid email or password")
	}

	userPass := []byte(Resp.Password)
	loginPass := []byte(login.Password)
	passErr := bcrypt.CompareHashAndPassword(userPass, loginPass)
	if passErr != nil {
		fmt.Println("Error in password", passErr)
		log.Println(passErr)
		return "", passErr
	}

	token, err := utils.GenerateToken(Resp.Email, Resp.Users_ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
