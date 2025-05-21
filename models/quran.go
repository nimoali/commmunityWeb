// models/quran.go
package models

type QuranVerse struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	TextAr    string `json:"text_ar"`
	TextEn    string `json:"text_en"`
	Reference string `json:"reference"`
	Tags      []string `json:"tags"`
	AudioURL  string `json:"audio_url"`
	Mood       string `bson:"mood" json:"mood"`
}



