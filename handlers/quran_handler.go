// handlers/session_handler.go
package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"newfolder/models"
	"newfolder/services"
	"time"

	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
	sessionService services.SessionService
	quranService   services.QuranService
}

func NewSessionHandler(ss services.SessionService, qs services.QuranService) *SessionHandler {
	return &SessionHandler{sessionService: ss, quranService: qs}
}

func (h *SessionHandler) StartSession(c *gin.Context) {
	var input struct {
		Goal     string `json:"goal"`
		Duration int    `json:"duration_minutes"`
		Mood     string `json:"mood"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	response, err := h.sessionService.StartSession(input.Goal, input.Duration, input.Mood)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start session"})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (h *SessionHandler) HandleAction(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Action string `json:"action"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	result := h.sessionService.HandleAction(id, input.Action)
	c.JSON(http.StatusOK, result)
}


func (h *QuranHandler) UploadVerse(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var verse models.QuranVerse
	err := json.NewDecoder(r.Body).Decode(&verse)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = h.QuranService.UploadVerses(ctx, &verse)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Verse uploaded successfully"})
}