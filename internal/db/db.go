package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jimitchavdadev/url-shortener/internal/config"
)

func NewDatabase(cfg *config.Config) (*sql.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Initialize schema
	if err := initializeSchema(db); err != nil {
		return nil, err
	}

	return db, nil
}

func initializeSchema(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS urls (
        id BIGINT AUTO_INCREMENT PRIMARY KEY,
        short_code VARCHAR(10) UNIQUE NOT NULL,
        original_url TEXT NOT NULL,
        click_count INT DEFAULT 0,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	_, err := db.Exec(query)
	if err != nil {
		log.Printf("Failed to initialize schema: %v", err)
		return err
	}
	return nil
}
