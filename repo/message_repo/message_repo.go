package message_repo

import "github.com/Zeeshanahmed-builds/ChatApp/models"

type MessageRepo interface {
	SaveMessage(message *models.Message) error
	GetMessages(senderID, receiverID int) ([]models.Message, error)
}
