// services/quran_service.go
package services

import (
	// "context"
	"context"
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

func (s *quranService) UploadVerse(ctx context.Context, verse *models.QuranVerse) error {
	return s.repo.InsertVerse(ctx, verse)
}


func (s *quranService) UploadVerses(ctx context.Context, verse *models.QuranVerse) error {
	return s.repo.UploadVerses(ctx, verse)
}