package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/note_project/storage/repo"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) repo.UserRepositoryStorage {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser()
