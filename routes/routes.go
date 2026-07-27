package routes

import (
	"chat-app/db"
	"chat-app/handler"
	"chat-app/repo"
	"chat-app/service/users_service/users_service_imp"
	"github.com/gin-gonic/gin"
)

func SetupRoutes( r *gin.Engine) {

	userRepo := repo.NewUsers(db.DB)


    authUsers :=users_service_imp.NewAuthUsers(users_service_imp.NewAuthUsersImp{
		Users: userRepo,
	})

    authHandler := handler.NewAuthHandler(
        authUsers,
    )

	r.POST("/signup", authHandler.SignUp)
	r.POST("/login",authHandler.Login)
}