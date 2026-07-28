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
	receiverID := c.GetInt("userID")
	fmt.Println("Receiver ID:", receiverID)
	message.ReceiverID = receiverID

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
<<<<<<< HEAD
	senderID := c.GetInt("user_id") // Logged-in user

	otherUserID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	messages, err := a.MessageService.GetReceivedMessages(senderID, otherUserID)
=======

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

>>>>>>> 1ca2fb4 (Refactor message handling: update message retrieval logic and modify routes; remove unused user handler)
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