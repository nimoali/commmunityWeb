// services/session_service.go
package services

import (
	"newfolder/models"
	repositories "newfolder/repository"
	"time"
)

type SessionService interface {
	StartSession(goal string, duration int, mood string) (map[string]interface{}, error)
	HandleAction(sessionID, action string) map[string]string
}

type sessionService struct {
	repo repositories.SessionRepository
	quran QuranService
}

func NewSessionService(repo repositories.SessionRepository) SessionService {
	return &sessionService{repo: repo}
}

func (s *sessionService) StartSession(goal string, duration int, mood string) (map[string]interface{}, error) {
	session := &models.Session{
		Goal:      goal,
		Duration:  duration,
		Mood:      mood,
		Status:    "running",
		CreatedAt: time.Now().Unix(),
	}
	sessionID, err := s.repo.CreateSession(session)
	if err != nil {
		return nil, err
	}

	verse, _ := s.quran.GetVerseByMood(mood)
	return map[string]interface{}{
		"session_id": sessionID,
		"verse":      verse,
	}, nil
}

func (s *sessionService) HandleAction(sessionID, action string) map[string]string {
	return map[string]string{
		"message": "Session " + action,
		"status":  action,
	}
}


