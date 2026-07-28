package message_service
import(
	"chat-app/models"
)


type MessageService interface {
	SaveMessage(message *models.Messages) error
	GetMessages(senderID, receiverID int) ([]models.Messages, error)

}