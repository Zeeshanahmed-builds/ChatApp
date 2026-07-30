package message_service

import (
	"chat-app/models"
)

type MessageService interface {
	SaveMessage(message *models.Message) error
	GetMessages(senderID, receiverID int) ([]models.Message, error)
}
