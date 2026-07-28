package users_repo_imp

import (
	"database/sql"
	"chat-app/repo/users_repo"
)

type Users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) users_repo.UsersRepo {
	return &Users{
		db: db,
	}

}
var _ users_repo.UsersRepo = (*Users)(nil)