package repository

import "database/sql"

type Notes struct {
	db *sql.DB
}

func NewNotes(db *sql.DB) *Notes {
	return &Notes{db: db}
}

func (n *Notes) GetNotes() {

}
