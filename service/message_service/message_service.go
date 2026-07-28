package message_service
import(
	"chat-app/models"
)


type MessageService interface {
	SaveMessage(message *models.Messages) error
	GetReceivedMessages(senderID, receiverID int) ([]models.Messages, error)
}