package main

import (
	"log"
	"net/http"

	"github.com/jimitchavdadev/url-shortener/internal/config"
	"github.com/jimitchavdadev/url-shortener/internal/db"
	"github.com/jimitchavdadev/url-shortener/internal/handlers"
	"github.com/jimitchavdadev/url-shortener/internal/repository"
	"github.com/jimitchavdadev/url-shortener/internal/routes"
	"github.com/jimitchavdadev/url-shortener/internal/services"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	dbConn, err := db.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	// Initialize repository, service, and handlers
	urlRepo := repository.NewURLRepository(dbConn)
	urlService := services.NewURLService(urlRepo)
	urlHandler := handlers.NewURLHandler(urlService)

	// Set up router
	router := routes.NewRouter(urlHandler)

	// Start server
	log.Printf("Server starting on :%s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
