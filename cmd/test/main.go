package main

import (
	"fmt"
	"github.com/fanfaronDo/code_education_api/internal/config"
	"github.com/fanfaronDo/code_education_api/internal/repository"

	_ "github.com/lib/pq"
	"log"
)

func main() {
	cfg := config.ConfigLoad()

	conn, err := repository.NewPostgres(cfg.Postgres)
	repo := repository.NewRepository(conn)
	user, err := repo.GetUser("Alex123", "123456")
	fmt.Println(user, err)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("User ID is", id)
}
