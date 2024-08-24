package repository

import "database/sql"

type AuthRepository interface {
	CreateUser()
	GetUser()
}

type NotesRepository interface {
	GetNotes()
}

type NoteRepository interface {
	CreateNote()
}

type Repository struct {
	Authorization
	NotesRepository
	NoteRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization:   NewAuthorization(db),
		NotesRepository: NewNotes(db),
		NoteRepository:  NewNote(db),
	}
}
