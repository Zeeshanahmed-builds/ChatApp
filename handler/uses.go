package handler
import (
	"fmt"
	"chat-app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (a AuthHandler)SignUp(c *gin.Context) {
	var users *models.Users
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

	
	// repo:=&repositories.SignRepo{}
	// service:=service.SingupService{
	// 	Repo:  repo,
	// }

	err = a.AuthUsers.SignUp(users)
	if err != nil {
		c.JSON(500, gin.H{"error": "email already exist"})
		return
	}
	// _,err = mail.ParseAddress(users.Email)
	// if err != nil{
	// 	c.JSON(500, gin.H{"error":"email is not valid "})
	// 	return
	// }

	c.JSON(200, gin.H{"message": "User signed up successfully"})
}



func (a AuthHandler)Login(c *gin.Context) {

	var login *models.Login

	if err :=c.ShouldBindJSON(&login) ;
	err != nil{
		c.JSON(400, gin.H{"error":"invalid data"})
		return
	}

	
	// repo:=&repositories.LoginRepo{}
	// service:=service.LoginService{
	// 	Repo:  repo,
	// }

	token,err := a.AuthUsers.Login(login)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "login successfully",
	"token":token})
}