// handlers/quran_handler.go
package handlers

import (
	"net/http"
	"newfolder/services"
	// "studyhard/services"
	"github.com/gin-gonic/gin"
)

type QuranHandler struct {
	service services.QuranService
}

func NewQuranHandler(service services.QuranService) *QuranHandler {
	return &QuranHandler{service: service}
}

func (h *QuranHandler) GetVerseByMood(c *gin.Context) {
	mood := c.Query("mood")
	verse, err := h.service.GetVerseByMood(mood)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Verse not found"})
		return
	}
	c.JSON(http.StatusOK, verse)
}