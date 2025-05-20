// models/session.go
package models

type Session struct {
	ID             string `json:"id" bson:"_id,omitempty"`
	Goal           string `json:"goal"`
	Duration       int    `json:"duration_minutes"`
	Mood           string `json:"mood"`
	Status         string `json:"status"`
	CreatedAt      int64  `json:"created_at" bson:"created_at"`
}