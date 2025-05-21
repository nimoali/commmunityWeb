package main

import (
	"context"
	"log"
	"newfolder/handlers"
	repositories "newfolder/repository"
	"newfolder/services"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}

	db := client.Database("studyhard")

	// Initialize Repositories
	sessionRepo := repositories.NewSessionRepository(db)
	quranRepo := repositories.NewQuranRepository(db)

	// Initialize Services
	sessionService := services.NewSessionService(sessionRepo)
	quranService := services.NewQuranService(quranRepo)

	// Initialize Handlers
	sessionHandler := handlers.NewSessionHandler(sessionService, quranService)
	quranHandler := handlers.NewQuranHandler(quranService)
	quranUploadRepo := repositories.NewQuranRepository(db)
	quranUplaodService := services.NewQuranService(quranUploadRepo)
	quranUploadHandler := handlers.NewSessionHandler(quranUplaodService)

	r := gin.Default()


	r.POST("/api/quran/upload", quranUploadHandler.UploadVerse)
	r.POST("/api/session/start", sessionHandler.StartSession)
	r.GET("/api/quran/verse", quranHandler.GetVerseByMood)
	r.POST("/api/session/:id/action", sessionHandler.HandleAction)

	r.Run(":8080")
}










