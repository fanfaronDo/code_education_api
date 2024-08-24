package repository

import "database/sql"

type Note struct {
	db *sql.DB
}

func NewNote(db *sql.DB) *Note {
	return &Note{db: db}
}

func (n *Note) CreateNote() {

}
