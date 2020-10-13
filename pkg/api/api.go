package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sirius-short-urls/configs"
	"sirius-short-urls/pkg/models"
	"sirius-short-urls/pkg/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// GenerateLink генерирует короткую ссылку и вставляет в базу
func GenerateLink(w http.ResponseWriter, r *http.Request) {
	// Парсим json в структуру
	var s models.ShortURL
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if s.OriginalURL == "" {
		http.Error(w, "", 400)
		return
	}

	// Генерим токен
	var tokenLength = configs.ProjectSettings.TokenLength
	token, err := utils.GenerateRandomString(tokenLength)
	if err != nil {
		log.Fatalln("Cant generate token: ", err)
		return
	}

	// Создаем в базе итоговый объект
	instance := models.GetShortURL(s.OriginalURL, token, s.IsReusable)
	_, err = configs.Collection.InsertOne(
		context.TODO(),
		instance,
	)
	if err != nil {
		log.Fatalln("Cant insert into collection: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Маршалим объект и возвращаем его
	js, err := json.Marshal(instance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Goto перенаправляет на оригинальный урл
func Goto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	token := params["token"]

	var s models.ShortURL

	err := configs.Collection.FindOne(
		context.TODO(),
		bson.M{"token": token},
	).Decode(&s)
	if err != nil {
		http.Error(w, "Token not found", 404)
		return
	}

	if !s.IsReusable {
		configs.Collection.DeleteOne(
			context.TODO(),
			bson.M{"_id": s.ID},
		)
	}
	http.Redirect(w, r, s.OriginalURL, 302)
}
