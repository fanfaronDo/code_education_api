package server

import (
	"context"
	"github.com/fanfaronDo/code_education_api/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (server *Server) Run(cfg config.HttpServer, httpHandler http.Handler) error {
	server.httpServer = &http.Server{
		Addr:           cfg.Address + ":" + cfg.Port,
		Handler:        httpHandler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    cfg.Timeout,
		WriteTimeout:   cfg.Timeout,
	}
	return server.httpServer.ListenAndServe()
}

func (server *Server) Stop(ctx context.Context) error {
	return server.httpServer.Shutdown(ctx)
}
