package shortner

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) ShortenURL(c *gin.Context) {
	var req struct {
		URL string `json:"url"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
		return
	}

	shortCode := h.svc.Shorten(req.URL)
	c.JSON(http.StatusOK, gin.H{
		"short_url": c.Request.Host + "/" + shortCode,
	})
}

func (h *Handler) ResolveURL(c *gin.Context) {
	short := c.Param("short")

	originalURL, found := h.svc.Resolve(short)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}