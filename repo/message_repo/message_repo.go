package message_repo
import "chat-app/models"

type MessageRepo interface {
	SaveMessage(message *models.Messages) error
<<<<<<< HEAD
<<<<<<< HEAD
	GetReceivedMessages(senderID, receiverID int) ([]models.Messages, error)
=======
	GetReceivedMessages(receiverID int) ([]models.Messages, error)
>>>>>>> b8f810e (Implement messaging functionality with database integration; add message handling and repository layers)
=======
	GetMessages(senderID, receiverID int) ([]models.Messages, error)
>>>>>>> 1ca2fb4 (Refactor message handling: update message retrieval logic and modify routes; remove unused user handler)
}