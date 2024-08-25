package service

import (
	"github.com/fanfaronDo/code_education_api/internal/domain"
	"github.com/fanfaronDo/code_education_api/internal/repository"
)

type Note struct {
	repo *repository.Repository
}

func NewNote(repo *repository.Repository) *Note {
	return &Note{repo: repo}
}

func (n *Note) CreateNote(userID int, note domain.Note) (int, error) {
	return n.repo.NoteRepository.CreateNote(userID, note)
}
