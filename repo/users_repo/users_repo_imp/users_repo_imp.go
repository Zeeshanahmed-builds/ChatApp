package users_repo_imp

import (
	"github.com/Zeeshanahmed-builds/ChatApp/repo/users_repo"
	"gorm.io/gorm"
)

type Users struct {
	db *gorm.DB
}

func NewUsers(db *gorm.DB) users_repo.UsersRepo {
	return &Users{db: db}

}

var _ users_repo.UsersRepo = (*Users)(nil)
