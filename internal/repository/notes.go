package repository

import (
	"database/sql"
	"github.com/fanfaronDo/code_education_api/internal/domain"
)

type Notes struct {
	db *sql.DB
}

func NewNotes(db *sql.DB) *Notes {
	return &Notes{db: db}
}

func (n *Notes) GetNotes(userID int) []domain.Note {
	var notes []domain.Note
	query := "SELECT title, description FROM notes WHERE user_id = $1"
	row, err := n.db.Query(query, userID)
	if err != nil {
		return nil
	}
	defer row.Close()
	for row.Next() {
		var note domain.Note
		err = row.Scan(&note.Title, &note.Description)
		if err != nil {
			return nil
		}
		notes = append(notes, note)
	}

	return notes
}
