package services

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/yourusername/URLShorten/internal/model"
	"github.com/yourusername/URLShorten/internal/repository"
)

type URLService struct {
	repo *repository.URLRepository
}

func New(repo *repository.URLRepository) *URLService {
	return &URLService{repo: repo}
}

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortCode() string {
	sb := strings.Builder{}
	for i := 0; i < 6; i++ {
		sb.WriteByte(alphabet[rand.Intn(len(alphabet))])
	}
	return sb.String()
}

func (service *URLService) ShortenURL(url string, expiresAt *time.Time) (*model.URL, error) {
	shortCode := generateShortCode()
	newURL := &model.URL{
		OriginalURL: url,
		ShortCode:   shortCode,
		ExpiresAt:   expiresAt,
	}
	err := service.repo.Save(newURL)
	if err != nil {
		return nil, err
	}
	return newURL, nil

}

func (service *URLService) GetURL(shortCode string) (*model.URL, error) {
	url, err := service.repo.FindByShortCode(shortCode)
	if err != nil {
		return nil, err
	}
	if url.ExpiresAt != nil && url.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("URL has expired")
	}
	err = service.repo.IncrementClickCount(shortCode)

	if err != nil {
		return nil, err
	}

	return url, nil

}

func (service *URLService) GetStats(shortCode string) (*model.URL, error) {
	url, err := service.repo.FindByShortCode(shortCode)
	if err != nil {
		return nil, err
	}
	return url, nil
}
