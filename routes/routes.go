package routes

import (
	"github.com/Zeeshanahmed-builds/ChatApp/handler"
	"github.com/Zeeshanahmed-builds/ChatApp/middleware"
	"github.com/Zeeshanahmed-builds/ChatApp/repo/message_repo/message_repo_impl"
	"github.com/Zeeshanahmed-builds/ChatApp/repo/users_repo/users_repo_imp"
	"github.com/Zeeshanahmed-builds/ChatApp/service/message_service/message_service_impl"
	"github.com/Zeeshanahmed-builds/ChatApp/service/users_service/users_service_impl"
	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, dbConn *gorm.DB, mqttClient paho.Client) {

	userRepo := users_repo_imp.NewUsers(dbConn)
	messageRepo := message_repo_imp.NewMessageRepo(dbConn)

	messageService := message_service_imp.NewMessageService(
		message_service_imp.NewMessageServiceImp{
			MessageRepo: messageRepo,
			MQTTClient:  mqttClient,
		},
	)

	authUsers := users_service_imp.NewAuthUsers(users_service_imp.NewAuthUsersImp{
		Users: userRepo,
	})

	authHandler := handler.NewAuthHandler(
		authUsers,
		messageService,
	)
	r.POST("/message/history", middleware.AuthMiddleware(), authHandler.GetMessages)
	r.POST("/message", middleware.AuthMiddleware(), authHandler.HandleSaveMessage)
	r.POST("/signup", authHandler.SignUp)
	r.POST("/login", authHandler.Login)
}
