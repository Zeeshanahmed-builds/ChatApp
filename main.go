package main

import (
	"chat-app/db"
	"chat-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	DB := db.ConnectDB()
	defer DB.Close()

	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run("localhost:8080")

}