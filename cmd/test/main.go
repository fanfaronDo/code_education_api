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
	fmt.Println(cfg)
	conn, err := repository.NewPostgres(cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Postgres: %s", conn)
}
