package message_repo
import "chat-app/models"

type MessageRepo interface {
	SaveMessage(message *models.Messages) error
	GetReceivedMessages(senderID, receiverID int) ([]models.Messages, error)
}