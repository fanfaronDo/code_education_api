package main

import (
	"fmt"
	"github.com/fanfaronDo/code_education_api/internal/config"
	"github.com/fanfaronDo/code_education_api/internal/repository"
	"github.com/fanfaronDo/code_education_api/internal/service"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.ConfigLoad()

	conn, err := repository.NewPostgres(cfg.Postgres)
	repo := repository.NewRepository(conn)
	serv := service.NewService(repo)

	//id, err := serv.CreateUser(domain.User{
	//	Name:     "Ann",
	//	Username: "Ann",
	//	Password: "1234",
	//})
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Printf("Created user with id %s\n", id)

	token, err := serv.AuthService.GenerateToken("Ann", "1234")
	fmt.Println(token)
	if err != nil {
		fmt.Println(err)
	}

	id, err := serv.AuthService.ParseToken(token)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(id)
}
