package main

import (
	"fmt"
	"github.com/fanfaronDo/code_education_api/internal/config"
	"github.com/fanfaronDo/code_education_api/internal/domain"
	"github.com/fanfaronDo/code_education_api/internal/repository"
	"github.com/fanfaronDo/code_education_api/internal/service"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.ConfigLoad()

	conn, err := repository.NewPostgres(cfg.Postgres)
	repo := repository.NewRepository(conn)
	serv := service.NewService(repo)

	id1 := 7

	//, err := serv.CreateUser(domain.User{
	//
	//	Name:     "Larisa",
	//	Username: "Beautiful",
	//	Password: "1234",
	//})
	//id2, err := serv.CreateUser(domain.User{
	//	Name:     "Vyacheslav",
	//	Username: "Handsome",
	//	Password: "1234",
	//})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Created user with id %s\n", id1)

	i, er := serv.NoteService.CreateNote(id1, domain.Note{
		Title:       "second",
		Description: "Test",
	})
	if er != nil {
		panic(er)
	}
	fmt.Printf("Created note with id %s, %s\n", id1, i)
	//
	//token, err := serv.AuthService.GenerateToken("Beautiful", "1234")

	//fmt.Println(token)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//id, err := serv.AuthService.ParseToken(token)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(id)
}
