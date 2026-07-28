package message_repo_imp
import (
	"chat-app/models"
)



func (m *MessageRepository) SaveMessage(message *models.Messages) error {

	_, err := m.db.Exec(
		"INSERT INTO messages(sender_id, receiver_id, message, is_read, created_at) VALUES ($1, $2, $3, $4, $5)",
		message.SenderID,
		message.ReceiverID,
		message.Message,
		message.IsRead,
		message.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}


<<<<<<< HEAD
func (m *MessageRepository) GetReceivedMessages(senderID, receiverID int) ([]models.Messages, error) {
=======
func (m *MessageRepository) GetReceivedMessages(receiverID int) ([]models.Messages, error) {
>>>>>>> b8f810e (Implement messaging functionality with database integration; add message handling and repository layers)

	
	rows, err := m.db.Query(`
		SELECT id, sender_id, receiver_id, message, is_read, created_at
		FROM messages
<<<<<<< HEAD
		WHERE (sender_id = $1 AND receiver_id = $2)
		   OR (sender_id = $2 AND receiver_id = $1)
		ORDER BY created_at ASC
	`, senderID, receiverID)
=======
		WHERE (receiver_id = $1)
		ORDER BY created_at ASC
	`, receiverID)
>>>>>>> b8f810e (Implement messaging functionality with database integration; add message handling and repository layers)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Messages

	for rows.Next() {
		var msg models.Messages

		err := rows.Scan(
			&msg.ID,
			&msg.SenderID,
			&msg.ReceiverID,
			&msg.Message,
			&msg.IsRead,
			&msg.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		messages = append(messages, msg)
	}
	// Check for errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}	