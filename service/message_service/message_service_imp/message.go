package message_service_imp
import(
	"chat-app/models"
)

func (m *MessageServiceImp) SaveMessage(message *models.Messages) error {
	err := m.messageRepo.SaveMessage(message)
	if err != nil {
		return err
	}
	return nil
}

<<<<<<< HEAD
func (m *MessageServiceImp) GetReceivedMessages(senderID, receiverID int) ([]models.Messages, error) {
	messages, err := m.messageRepo.GetReceivedMessages(senderID, receiverID)
=======
func (m *MessageServiceImp) GetReceivedMessages(receiverID int) ([]models.Messages, error) {
	messages, err := m.messageRepo.GetReceivedMessages(receiverID)
>>>>>>> b8f810e (Implement messaging functionality with database integration; add message handling and repository layers)
	if err != nil {
		return nil, err
	}
	return messages, nil
}