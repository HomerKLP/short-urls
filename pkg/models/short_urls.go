package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ShortURL - модель коротких ссылок
type ShortURL struct {
	ID          primitive.ObjectID `bson:"_id"`
	OriginalURL string             `json:"original_url" bson:"original_url"`
	Token       string             `json:"token" bson:"token"`
	IsReusable  bool               `json:"is_reusable" bson:"is_reusable"`
	CreatedAt   int64              `json:"created_at" bson:"created_at"`
}

// GetShortURL реализует объект модели ShortURL
func GetShortURL(originalURL string, token string, isReusable bool) *ShortURL {
	var instance = ShortURL{
		ID:          primitive.NewObjectID(),
		OriginalURL: originalURL,
		Token:       token,
		IsReusable:  isReusable,
		CreatedAt:   time.Now().Unix(),
	}
	return &instance
}
