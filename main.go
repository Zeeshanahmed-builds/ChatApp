package main

import (
	"chat-app/db"
	"chat-app/mqtt"
	"chat-app/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	mqttClient, err := mqtt.Connect()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to EMQX")

	if err := mqtt.Subscribe(mqttClient); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	routes.SetupRoutes(router, db, mqttClient)
	err = router.Run("localhost:8080")
	if err != nil {
	log.Fatal(err)
	}
}
