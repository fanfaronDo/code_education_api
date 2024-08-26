package main

import (
	"fmt"
	"github.com/fanfaronDo/code_education_api/internal/config"
	"github.com/fanfaronDo/code_education_api/internal/handler"
	"github.com/fanfaronDo/code_education_api/internal/repository"
	"github.com/fanfaronDo/code_education_api/internal/server"
	"github.com/fanfaronDo/code_education_api/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.ConfigLoad()

	server := &server.Server{}
	conn, err := repository.NewPostgres(cfg.Postgres)
	fmt.Println(cfg, conn)
	if err != nil {
		log.Printf("Database connection error: %s\n", err)
		return
	}
	repo := repository.NewRepository(conn)
	service := service.NewService(repo)
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
