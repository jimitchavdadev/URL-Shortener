package services

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/jimitchavdadev/url-shortener/internal/models"
	"github.com/jimitchavdadev/url-shortener/internal/repository"
)

type URLService struct {
	repo *repository.URLRepository
}

func NewURLService(repo *repository.URLRepository) *URLService {
	return &URLService{repo: repo}
}

func (s *URLService) ShortenURL(originalURL string) (*models.URL, error) {
	shortCode := generateShortCode()
	url := &models.URL{
		OriginalURL: originalURL,
		ShortCode:   shortCode,
	}

	if err := s.repo.Save(url); err != nil {
		return nil, err
	}
	return url, nil
}

func (s *URLService) GetOriginalURL(shortCode string) (*models.URL, error) {
	url, err := s.repo.FindByShortCode(shortCode)
	if err != nil {
		return nil, err
	}

	if err := s.repo.IncrementClickCount(shortCode); err != nil {
		return nil, err
	}
	return url, nil
}

func generateShortCode() string {
	b := make([]byte, 6)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:8]
}
