package routes

import (
	"database/sql"
	"chat-app/handler"
	"chat-app/middleware"
	"chat-app/repo/message_repo/message_repo_impl"
	"chat-app/repo/users_repo/users_repo_imp"
	"chat-app/service/message_service/message_service_impl"
	"chat-app/service/users_service/users_service_impl"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
)

func SetupRoutes( r *gin.Engine, dbConn *sql.DB, mqttClient paho.Client) {

	userRepo := users_repo_imp.NewUsers(dbConn)
	messageRepo := message_repo_imp.NewMessageRepo(dbConn)

	messageService := message_service_imp.NewMessageService(
    message_service_imp.NewMessageServiceImp{
        MessageRepo: messageRepo,
		MQTTClient:  mqttClient,
    },
)

    authUsers :=users_service_imp.NewAuthUsers(users_service_imp.NewAuthUsersImp{
		Users: userRepo,
	})

    authHandler := handler.NewAuthHandler(
        authUsers,
        messageService,
    )
	r.POST("/message/history", middleware.AuthMiddleware(), authHandler.GetMessages)
	r.POST("/message",middleware.AuthMiddleware(), authHandler.HandleSaveMessage)
	r.POST("/signup", authHandler.SignUp)
	r.POST("/login",authHandler.Login)
}