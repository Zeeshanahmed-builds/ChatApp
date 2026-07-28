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

func (m *MessageServiceImp) GetMessages(senderID, receiverID int) ([]models.Messages, error) {
	messages, err := m.messageRepo.GetMessages(senderID, receiverID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}