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


func (m *MessageRepository) GetMessages(senderID, receiverID int) ([]models.Messages, error) {

	rows, err := m.db.Query(`
		SELECT
			id,
			sender_id,
			receiver_id,
			message,
			is_read,
			created_at
		FROM messages

		WHERE (sender_id = $1 AND receiver_id = $2)
		   OR (sender_id = $2 AND receiver_id = $1)
		
	`, senderID, receiverID)

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