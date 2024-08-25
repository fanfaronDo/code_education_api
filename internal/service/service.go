package service

import (
	"github.com/fanfaronDo/code_education_api/internal/domain"
	"github.com/fanfaronDo/code_education_api/internal/repository"
)

type AuthService interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type NoteService interface {
	CreateNote(userID int, note domain.Note) (int, error)
}

type NotesService interface {
	GetNotes(userID int) []domain.Note
}

type Service struct {
	AuthService
	NoteService
	NotesService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthService:  NewAuthorization(repo),
		NoteService:  NewNote(repo),
		NotesService: NewNotes(repo),
	}
}
