package repository

import (
	"database/sql"
	"github.com/fanfaronDo/code_education_api/internal/domain"
)

type AuthRepository interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type NotesRepository interface {
	GetNotes()
}

type NoteRepository interface {
	CreateNote()
}

type Repository struct {
	AuthRepository
	NotesRepository
	NoteRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		AuthRepository:  NewAuthorization(db),
		NotesRepository: NewNotes(db),
		NoteRepository:  NewNote(db),
	}
}
