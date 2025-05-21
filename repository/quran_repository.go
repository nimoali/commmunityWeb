// repositories/quran_repository.go
package repositories

import (
	"context"
	"newfolder/models"

	// "studyhard/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuranRepository interface {
	FindVerseByTag(tag string) (*models.QuranVerse, error)
	InsertVerse(ctx context.Context, verse *models.QuranVerse) error
	UploadVerses(ctx context.Context, verse *models.QuranVerse) error
}

// func (q QuranRepository) UploadVerses(ctx context.Context, verse *models.QuranVerse) error {
// 	panic("unimplemented")
// }

func (r *quranRepo) InsertVerse(ctx context.Context, verse *models.QuranVerse) error {
	_, err := r.collection.InsertOne(ctx, verse)
	return err
}

type quranRepo struct {
	collection *mongo.Collection
}

func NewQuranRepository(db *mongo.Database) QuranRepository {
	return &quranRepo{collection: db.Collection("quran_verses")}
}

func (r *quranRepo) FindVerseByTag(tag string) (*models.QuranVerse, error) {
	filter := bson.M{"tags": tag}
	var verse models.QuranVerse
	err := r.collection.FindOne(context.TODO(), filter).Decode(&verse)
	if err != nil {
		return nil, err
	}
	return &verse, nil
}

func (r *quranRepo) UploadVerses(ctx context.Context, verse *models.QuranVerse) error {
	_, err := r.collection.InsertOne(ctx, verse)
	return err
}
