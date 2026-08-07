package message_service_imp

import (
	"github.com/Zeeshanahmed-builds/ChatApp/models"
	"encoding/json"
	"fmt"
	"time"
)

func (m *MessageServiceImp) SaveMessage(message *models.Message) error {
	err := m.messageRepo.SaveMessage(message)
	if err != nil {
		return err
	}

	topic := fmt.Sprintf("chat/%d/inbox", message.ReceiverID)

	payload, err := json.Marshal(map[string]interface{}{
		"sender_id": message.SenderID,
		"message":   message.Message,
		"timestamp": time.Now().Unix(),
	})
	if err != nil {
		return err
	}

	token := m.mqttClient.Publish(
		topic,
		1,
		false,
		payload,
	)

	token.Wait()
	return token.Error()
}

func (m *MessageServiceImp) GetMessages(senderID, receiverID int) ([]models.Message, error) {
	messages, err := m.messageRepo.GetMessages(senderID, receiverID)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
