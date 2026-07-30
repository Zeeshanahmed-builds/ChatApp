package handler

import (
	"chat-app/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (a AuthHandler) SignUp(c *gin.Context) {
	var users *models.User
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	validate := validator.New()
	err := validate.Struct(users)
	if err != nil {
		fmt.Println("Validation failed:", err.Error())
		c.JSON(400, gin.H{"error": err})
		return
	}

	fmt.Println(users, users)

	err = a.AuthUsers.SignUp(users)
	if err != nil {
		c.JSON(500, gin.H{"error": "email already exist"})
		return
	}

	c.JSON(200, gin.H{"message": "User signed up successfully"})
}

func (a AuthHandler) Login(c *gin.Context) {

	var login *models.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := a.AuthUsers.Login(login)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "login successfully",
		"token": token})
}
