package main

import (
	"log"
	"net/http"
	"server/internal/repo"
)

func main() {

	repos, err := repo.InitRepo()
	if err != nil {
		log.Fatalf("Failed to initialize db: %v", err)
	}

	s := service.NewService(repos)
	routes := handlers.NewHandlers(s)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: routes.InitHandlers(),
	}

	log.Println("Starting server on localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
