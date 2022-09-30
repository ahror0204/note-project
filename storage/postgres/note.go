package postgres

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/note_project/pkg/structures"
	repo "github.com/note_project/storage/repo"
)

type noteRepo struct {
	db *sqlx.DB
}

func NewnoteRepo(db *sqlx.DB) repo.NoteRepositoryStorage {
	return &noteRepo{
		db: db,
	}
}

func (r *noteRepo) CreateNote(note *structures.NoteStruct) (*structures.NoteStruct, error) {
	query := `INSERT INTO note_table(id, title, body, exp_time, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id, title, body`
	var rnote structures.NoteStruct
	createdAt := time.Now()

	ID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(query, ID, note.Title, note.Body, note.ExpTime, createdAt).Scan(
		&rnote.ID,
		&rnote.Title,
		&rnote.Body,
	)
	if err != nil {
		return nil, err
	}

	return &rnote, nil
}

func (r *noteRepo) UpdateNote(note *structures.NoteStruct) (*structures.NoteStruct, error) {
	var rnote structures.NoteStruct
	query := `UPDATE note_table SET title = $1, body = $2, exp_time = $3, updated_at = $4 WHERE id = $5 AND deleted_at IS NULL
		RETURNING id, title, body, exp_time, updated_at`

	createdAT := time.Now()

	err := r.db.QueryRow(query, note.Title, note.Body, note.ExpTime, createdAT, note.ID).Scan(
		&rnote.ID,
		&rnote.Title,
		&rnote.Body,
		&rnote.ExpTime,
		&rnote.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &rnote, nil
}

func (r *noteRepo) DeleteNote(ID string) error {

	query := `UPDATE note_table SET deleted_at = $1 WHERE id = $2 AND deleted_at is NULL`

	deletedAT := time.Now()

	_, err := r.db.Query(query, deletedAT, ID)

	if err != nil {
		return err
	}

	return nil
}
