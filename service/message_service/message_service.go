package message_service

import (
	"github.com/Zeeshanahmed-builds/ChatApp/models"
)

type MessageService interface {
	SaveMessage(message *models.Message) error
	GetMessages(senderID, receiverID int) ([]models.Message, error)
}
