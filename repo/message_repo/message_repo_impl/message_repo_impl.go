package message_repo_imp

import (
	"chat-app/repo/message_repo"
	"database/sql"
)

type MessageRepository struct {
	db *sql.DB
}

func NewMessageRepo(db *sql.DB) message_repo.MessageRepo {
	return &MessageRepository{db: db}
}

var _ message_repo.MessageRepo = (*MessageRepository)(nil)