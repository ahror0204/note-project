package postgres

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/note_project/pkg/structures"
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

func (r *userRepo) CreateUser(user *structures.UserStruct) (*structures.UserStruct, error) {
	createdTime := time.Now()
	ID, err := uuid.NewV4()
	var ruser structures.UserStruct

	if err != nil {
		return nil, err
	}
	query := `INSERT INTO users(id, first_name, last_name, email, created_at) 
			  VALUES ($1,$2,$3,$4,$5)
			  RETURNING id, first_name, last_name, email, created_at`

	err = r.db.QueryRow(query, ID, user.FirstName, user.LastName, user.Email, createdTime).Scan(
		&ruser.ID,
		&ruser.FirstName,
		&ruser.LastName,
		&ruser.Email,
		&ruser.CreatedAt,
	)

	return &ruser, nil
}

func (r *userRepo) UpdateUser(user *structures.UserStruct) (error) {
	updatedAt := time.Now()
	query := `UPDATE users SET first_name = $1, last_name = $2, email = $3, updated_at = $4 WHERE id = $5`

	_, err := r.db.Query(query, user.FirstName, user.LastName, user.Email, updatedAt, user.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) DeleteUser(ID string) error {
	query := `UPDATE users SET deleted_at = $1 WHERE id = $2`

	deletedAt := time.Now()
	_, err := r.db.Query(query, deletedAt, ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) CheckField(req structures.UserCheckRequest) (bool, error) {
	
	var count int
	if req.Field == "username" {
		err := r.db.QueryRow("SELECT COUNT(1) FROM users WHERE user_name = $1 AND deleted_at = NULL", req.Value).Scan(&count)
		if err != nil {
			return false, err
		}
	} else if req.Field == "email" {
		err := r.db.QueryRow("SELECT COUNT(1) FROM users WHERE email = $1 AND deleted_at = NULL", req.Value).Scan(&count)
		if err != nil {
			return false, err
		}
	} else {
		err := errors.New("CheckField func error!")
			return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}