package message_repo
import "chat-app/models"

type MessageRepo interface {
	SaveMessage(message *models.Message) error
	GetMessages(senderID, receiverID int) ([]models.Message, error)

}