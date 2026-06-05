package model

import "time"

type URL struct {
	ID          int        `db:"id"          json:"id"`
	OriginalURL string     `db:"original_url" json:"original_url"`
	ShortCode   string     `db:"short_code"   json:"short_code"`
	CreatedAt   time.Time  `db:"created_at"   json:"created_at"`
	ClickCount  int        `db:"click_count"  json:"click_count"`
	ExpiresAt   *time.Time `db:"expires_at"   json:"expires_at,omitempty"`
}
