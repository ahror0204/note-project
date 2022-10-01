package repo

import "github.com/note_project/pkg/structures"

type NoteRepositoryStorage interface {
	CreateNote(*structures.NoteStruct) (*structures.NoteStruct, error)
	UpdateNote(*structures.NoteStruct) (*structures.NoteStruct, error)
	DeleteNote(Id string) error
}

type UserRepositoryStorage interface {
	CreateUser(*structures.UserStruct) (*structures.UserStruct, error)
	UpdateUser(*structures.UserStruct) (error)
	DeleteUser(id string) (error)
}
