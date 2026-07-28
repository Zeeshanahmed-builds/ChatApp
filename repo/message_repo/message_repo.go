package message_repo
import "chat-app/models"

type MessageRepo interface {
	SaveMessage(message *models.Messages) error
	GetMessages(senderID, receiverID int) ([]models.Messages, error)

}