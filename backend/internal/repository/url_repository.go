package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/yourusername/URLShorten/internal/model"
)

type URLRepository struct {
	db  *sqlx.DB
	rdb *redis.Client
}

func New(db *sqlx.DB, rdb *redis.Client) *URLRepository {
	return &URLRepository{db: db, rdb: rdb}
}
func (r *URLRepository) Save(url *model.URL) error {
	query := `INSERT INTO urls(original_url,short_code,expires_at)
			VALUES (:original_url,:short_code,:expires_at)
	`
	_, err := r.db.NamedExec(query, url)
	return err
}

func (r *URLRepository) FindByShortCode(shortCode string) (*model.URL, error) {
	url := &model.URL{}
	val, err := r.rdb.Get(context.Background(), shortCode).Result()
	if err == nil {
		json.Unmarshal([]byte(val), url)
		return url, nil
	}
	query := `SELECT * FROM urls WHERE short_code = $1`
	err = r.db.Get(url, query, shortCode)
	if err != nil {
		return nil, err
	}
	jsonBytes, _ := json.Marshal(url)
	r.rdb.Set(context.Background(), shortCode, jsonBytes, time.Hour)
	return url, err
}

func (r *URLRepository) IncrementClickCount(shortCode string) error {
	query := `UPDATE urls SET click_count = click_count + 1 WHERE short_code =$1`
	_, err := r.db.Exec(query, shortCode)
	return err
}
