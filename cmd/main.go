package main

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "SERVER STATUS", log.LstdFlags)
	srv := server.NewServer(logger)
	logger.Println("starting server 8080")
	if err := srv.HTTP.ListenAndServe(); err != nil {
		logger.Fatal("server failed to start: ", err)
	}
}
