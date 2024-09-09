package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fanfaronDo/code_education_api/internal/config"
	"github.com/fanfaronDo/code_education_api/internal/domain"
	"github.com/fanfaronDo/code_education_api/internal/handler"
	"github.com/fanfaronDo/code_education_api/internal/repository"
	"github.com/fanfaronDo/code_education_api/internal/server"
	"github.com/fanfaronDo/code_education_api/internal/service"
)

func main() {
	cfg := config.ConfigLoad()
	server := make(server.Server)
	conn, err := repository.NewPostgres(cfg.Postgres)
	if err != nil {
		log.Printf("Database connection error: %s\n", err)
		return
	}
	repo := repository.NewRepository(conn)
	service := service.NewService(repo)

	user1 := domain.User{
		Name:     "John Doe",
		Username: "john_doe",
		Password: "password123",
	}

	user2 := domain.User{
		Name:     "Jane Smith",
		Username: "jane_smith",
		Password: "secret_pass",
	}

	service.AuthService.CreateUser(user1)
	service.AuthService.CreateUser(user2)

	h := handler.NewHandler(service)
	route := h.InitRoutes()
	go func() {
		if err := server.Run(cfg.HttpServer, route); err != nil {
			log.Printf("Error starting server: %v\n", err)
			return
		}
	}()
	defer server.Stop(nil)

	log.Printf("Server started on %s\n", "http://"+cfg.HttpServer.Address+":"+cfg.HttpServer.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Printf("Shutting down server...\n")
}
