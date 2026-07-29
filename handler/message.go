package handler

import (
	"fmt"
	"chat-app/models"
	"github.com/gin-gonic/gin"
)
func (a *AuthHandler) HandleSaveMessage(c *gin.Context) {

	var message models.Messages

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Get sender ID from JWT
	senderID := c.GetInt("userID")
	fmt.Println("Receiver ID:", senderID)
	message.SenderID = senderID

	// New messages are always unread
	// message.IsRead = false

	if err := a.MessageService.SaveMessage(&message); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Message sent successfully",
	})
}

func (a *AuthHandler) GetMessages(c *gin.Context) {

	senderID := c.GetInt("userID")

	var req struct {
		ReceiverID int `json:"receiver_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	messages, err := a.MessageService.GetMessages(senderID, req.ReceiverID)

	if err != nil {

		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"messages": messages,
	})
}