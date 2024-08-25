package service

import (
	"github.com/fanfaronDo/code_education_api/internal/domain"
	"github.com/fanfaronDo/code_education_api/internal/repository"
)

type Notes struct {
	repo *repository.Repository
}

func NewNotes(repo *repository.Repository) *Notes {
	return &Notes{repo: repo}
}

func (n *Notes) GetNotes(userID int) []domain.Note {
	return n.repo.GetNotes(userID)
}
