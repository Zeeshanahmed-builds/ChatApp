package handler

import (
	"fmt"
	"net/http"
	"strconv"
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
	fmt.Println("Sender ID:", senderID)
	message.SenderID = senderID

	// New messages are always unread
	message.IsRead = false

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
	senderID := c.GetInt("user_id") // Logged-in user

	otherUserID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	messages, err := a.MessageService.GetReceivedMessages(senderID, otherUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": messages})
}