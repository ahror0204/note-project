package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/note_project/storage/postgres"
	"github.com/note_project/storage/repo"
)

type storagePg struct {
	db       *sqlx.DB
	NoteRepo repo.NoteRepositoryStorage
}

func NewPostgrepg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		NoteRepo: postgres.NewnoteRepo(db),
	}
}
