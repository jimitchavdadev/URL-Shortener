package models

import "time"

type URL struct {
	ID          int64
	ShortCode   string
	OriginalURL string
	ClickCount  int
	CreatedAt   time.Time
}
