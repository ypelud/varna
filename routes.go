package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes ...
type Routes []Route

// NewRouter ...
func NewRouter(globalConfiguration *GlobalConfiguration) *mux.Router {

	handlers := NewHandlers(globalConfiguration)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", handlers.HealthPage).Methods("GET")
	router.HandleFunc("/webhook", handlers.DockerWebhook).Methods("POST")

	return router
}
