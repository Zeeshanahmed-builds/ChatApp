package message_repo_imp

import (
	"chat-app/repo/message_repo"
	"gorm.io/gorm"

)

type MessageRepository struct {
	// db *sql.DB
	db *gorm.DB

}

func NewMessageRepo(db *gorm.DB) message_repo.MessageRepo {
	return &MessageRepository{db: db}
}

var _ message_repo.MessageRepo = (*MessageRepository)(nil)