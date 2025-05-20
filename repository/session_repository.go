// repositories/session_repository.go
package repositories

import (
	"context"
	"newfolder/models"
	// "studyhard/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionRepository interface {
	CreateSession(session *models.Session) (string, error)
}

type sessionRepo struct {
	collection *mongo.Collection
}

func NewSessionRepository(db *mongo.Database) SessionRepository {
	return &sessionRepo{collection: db.Collection("sessions")}
}

func (r *sessionRepo) CreateSession(session *models.Session) (string, error) {
	res, err := r.collection.InsertOne(context.TODO(), session)
	if err != nil {
		return "", err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}