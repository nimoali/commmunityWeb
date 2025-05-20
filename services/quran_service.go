// services/quran_service.go
package services

import (
	"newfolder/models"
	repositories "newfolder/repository"
)

type QuranService interface {
	GetVerseByMood(mood string) (*models.QuranVerse, error)
}

type quranService struct {
	repo repositories.QuranRepository
}

func NewQuranService(repo repositories.QuranRepository) QuranService {
	return &quranService{repo: repo}
}

func (s *quranService) GetVerseByMood(mood string) (*models.QuranVerse, error) {
	return s.repo.FindVerseByTag(mood)
}