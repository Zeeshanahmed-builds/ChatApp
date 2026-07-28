package message_service
import(
	"chat-app/models"
)


type MessageService interface {
	SaveMessage(message *models.Messages) error
<<<<<<< HEAD
	GetReceivedMessages(senderID, receiverID int) ([]models.Messages, error)
=======
	GetReceivedMessages(receiverID int) ([]models.Messages, error)
>>>>>>> b8f810e (Implement messaging functionality with database integration; add message handling and repository layers)
}