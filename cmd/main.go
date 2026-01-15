package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	// логгер
	logger := log.New(os.Stdout, "SERVER: ", log.Ldate|log.Ltime)

	// сервер
	srv := server.New(logger)

	// запуск сервера
	logger.Println("Starting server on:8080")
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatalf("Server failed: %v", err)
	}
}
