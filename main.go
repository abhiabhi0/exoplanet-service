package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func generateID() string {
	return uuid.New().String()
}

func main() {
	storage := NewStorage()

	r := mux.NewRouter()
	r.HandleFunc("/exoplanets", addExoplanetHandler(storage)).Methods("POST")
	r.HandleFunc("/exoplanets", listExoplanetsHandler(storage)).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", getExoplanetByIDHandler(storage)).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", updateExoplanetHandler(storage)).Methods("PUT")
	r.HandleFunc("/exoplanets/{id}", deleteExoplanetHandler(storage)).Methods("DELETE")
	r.HandleFunc("/exoplanets/{id}/fuel", calculateFuelHandler(storage)).Methods("GET")

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
