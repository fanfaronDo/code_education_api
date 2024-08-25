package repository

import (
	"database/sql"
	"github.com/fanfaronDo/code_education_api/internal/domain"
)

type Note struct {
	db *sql.DB
}

func NewNote(db *sql.DB) *Note {
	return &Note{db: db}
}

func (n *Note) CreateNote(userID int, note domain.Note) (int, error) {
	var id int
	query := "INSERT INTO notes (title, description, user_id) VALUES ($1, $2, $3) RETURNING note_id"
	rew := n.db.QueryRow(query, note.Title, note.Description, userID)
	if err := rew.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
