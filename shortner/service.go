package shortner

import (
	"math/rand"
	"time"
)

type Service interface {
	Shorten(originalURL string) string
	Resolve(shortCode string) (string, bool)
}

type service struct {
	urls map[string]string
}

func NewService() Service {
	return &service{
		urls: make(map[string]string),
	}
}

func (s *service) Shorten(originalURL string) string {
	rand.Seed(time.Now().UnixNano())
	shortCode := RandString(6)
	s.urls[shortCode] = originalURL
	return shortCode
}

func (s *service) Resolve(shortCode string) (string, bool) {
	url, found := s.urls[shortCode]
	return url, found
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}