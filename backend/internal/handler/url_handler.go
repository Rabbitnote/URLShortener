package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/URLShorten/internal/services"
)

type URLHandler struct {
	service *services.URLService
}
type ShortenRequest struct {
	OriginalURL string     `json:"original_url"`
	ExpiresAt   *time.Time `json:"expires_at"`
}

func New(s *services.URLService) *URLHandler {
	return &URLHandler{service: s}
}
func (h *URLHandler) ShortenURL(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	url, err := h.service.ShortenURL(req.OriginalURL, req.ExpiresAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"short_url": url.ShortCode})

}
func (h *URLHandler) GetURLHandler(c *gin.Context) {
	shortCode := c.Param("code")
	url, err := h.service.GetURL(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}
func (h *URLHandler) GetStatsHandler(c *gin.Context) {
	shortCode := c.Param("code")
	url, err := h.service.GetStats(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, url)
}
