package controllers

import (
	"sirius-short-urls/pkg/api"

	"github.com/gorilla/mux"
)

// Router - Основной рутер сервиса
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/generate-link/", api.GenerateLink).Methods("POST")
	router.HandleFunc("/{token}", api.Goto).Methods("GET")
	return router
}
