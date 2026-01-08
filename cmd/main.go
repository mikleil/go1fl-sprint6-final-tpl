package main

import (
	"log"
	"os"

	"go1fl-sprint6-final-tpl/internal/handlers"
	"go1fl-sprint6-final-tpl/internal/server"
	"go1fl-sprint6-final-tpl/internal/service"
	"go1fl-sprint6-final-tpl/pkg/morse"
)

func main() {
	log := log.New(os.Stdout, "", 0)

	converter := morse.NewConverter(morse.DefaultMorse)
	mainService := service.NewService(converter)
	mainHandler := handlers.New(log, mainService)
	srv := server.New(log, "localhost", "8080", 10, 5, 15, mainHandler)
	srv.Start()
}
