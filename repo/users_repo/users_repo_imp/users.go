package users_repo_imp

import (
	"github.com/Zeeshanahmed-builds/ChatApp/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func (u *Users) SignUp(user *models.User) error {

	log.Println("Signing up user:", user.Name, user.Email)

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return err
	}

	user.Password = string(hashPassword)

	err = u.db.Create(user).Error
	if err != nil {
		return err
	}
	log.Println("User signed up successfully:", user.Name)
	return nil

}

func (u *Users) Login(email string) (*models.User, error) {
	var user models.User

	err := u.db.
		Where("email = ?", email).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, err
}
