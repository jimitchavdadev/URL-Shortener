package repository

import (
	"database/sql"
	"github.com/jimitchavdadev/url-shortener/internal/models"
)

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Save(url *models.URL) error {
	query := `INSERT INTO urls (short_code, original_url) VALUES (?, ?)`
	result, err := r.db.Exec(query, url.ShortCode, url.OriginalURL)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	url.ID = id
	return nil
}

func (r *URLRepository) FindByShortCode(shortCode string) (*models.URL, error) {
	url := &models.URL{}
	query := `SELECT id, short_code, original_url, click_count, created_at FROM urls WHERE short_code = ?`
	err := r.db.QueryRow(query, shortCode).Scan(&url.ID, &url.ShortCode, &url.OriginalURL, &url.ClickCount, &url.CreatedAt)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (r *URLRepository) IncrementClickCount(shortCode string) error {
	query := `UPDATE urls SET click_count = click_count + 1 WHERE short_code = ?`
	_, err := r.db.Exec(query, shortCode)
	return err
}
