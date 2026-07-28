package main

import (
	"log"
	"chat-app/db"
	"chat-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	dbConn, err := db.ConnectDB()
if err != nil {
	log.Fatal(err)
}
defer dbConn.Close()

	router := gin.Default()
	routes.SetupRoutes(router,dbConn)
	router.Run("localhost:8080")

}