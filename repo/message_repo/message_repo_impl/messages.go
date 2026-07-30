package message_repo_imp

import (
	"chat-app/models"
)

func (m *MessageRepository) SaveMessage(message *models.Message) error {

	err := m.db.Create(message).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *MessageRepository) GetMessages(senderID, receiverID int) ([]models.Message, error) {

	var messages []models.Message

	err := m.db.
		Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
			senderID, receiverID, receiverID, senderID).
		Order("created_at asc").
		Find(&messages).Error

	return messages, err
}
