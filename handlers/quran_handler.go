// handlers/session_handler.go
package handlers

import (
	"net/http"
	"newfolder/services"

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
