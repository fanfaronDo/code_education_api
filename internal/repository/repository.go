package repository

import "database/sql"

type Authorization interface {
	CreateUser()
	GetUser()
}

type Notes interface {
	GetNotes()
}

type Note interface {
	CreateNote()
}

type Repository struct {
	Authorization
	Notes
	Note
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthorization(db),
		Notes:         NewNotes(db),
		Note:          NewNote(db),
	}
}
