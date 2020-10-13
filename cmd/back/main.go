package main

import (
	"log"
	"net/http"
	"sirius-short-urls/configs"
	"sirius-short-urls/pkg/controllers"
)

func main() {
	log.Println("Sirius-short-urls server started...")
	// Настройки проекта
	var settings = configs.Settings()

	// Подключаемся к MongoDB
	err := configs.ConnectDB(settings.DBUri)
	if err != nil {
		log.Fatalln("Error while connecting to DB: ", err)
	}

	// Настраиваем сервер
	server := &http.Server{
		Addr:    "0.0.0.0:" + settings.AppPort,
		Handler: controllers.Router(),
	}
	serverError := server.ListenAndServe()
	if serverError != nil {
		log.Fatal("Error while running server: ", serverError)
	}
}
