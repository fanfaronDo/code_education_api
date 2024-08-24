package repository

import (
	"database/sql"
	"github.com/fanfaronDo/code_education_api/internal/domain"
)

type Authorization struct {
	db *sql.DB
}

func NewAuthorization(db *sql.DB) *Authorization {
	return &Authorization{db: db}
}

func CreateUser(user domain.User) (int, error) {
	var id int
	query := "INSERT INTO users (user_id, ) " +
	return 0, nil
}

func GetUser() (int, error) {
	return 0, nil
}
