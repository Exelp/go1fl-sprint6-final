package server

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Logger *log.Logger
	Http   *http.Server
}

func NewServer(loger *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.ReturnIndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     loger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &Server{
		Logger: loger,
		Http:   httpServer,
	}
}
