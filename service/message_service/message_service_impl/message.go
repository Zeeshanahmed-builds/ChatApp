package message_service_imp
import(
	"fmt"
	"chat-app/models"
)

func (m *MessageServiceImp) SaveMessage(message *models.Messages) error {
	err := m.messageRepo.SaveMessage(message)
	if err != nil {
		return err
	}

	
	topic := fmt.Sprintf("chat/%d/inbox", message.ReceiverID)

	token := m.mqttClient.Publish(
		topic,
		1,     
		false,
		message.Message,
	)

    token.Wait()
    return token.Error()
}


func (m *MessageServiceImp) GetMessages(senderID, receiverID int) ([]models.Messages, error) {
	messages, err := m.messageRepo.GetMessages(senderID, receiverID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}