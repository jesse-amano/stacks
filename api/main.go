package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// See all books
	router.Handle("/api/inventory/", TODO).Methods(http.MethodGet)
	// Define a new book
	router.Handle("/api/inventory/", TODO).Methods(http.MethodPost)
	// View book description
	router.Handle("/api/inventory/{:isbn}", TODO).Methods(http.MethodGet)
	// Modify book description
	router.Handle("/api/inventory/{:isbn}", TODO).Methods(http.MethodPut)
	// Remove book description
	router.Handle("/api/inventory/{:isbn}", TODO).Methods(http.MethodDelete)
	// Count copies of book
	router.Handle("/api/inventory/{:isbn}/copies", TODO).Methods(http.MethodGet)
	// Add copy of book
	router.Handle("/api/inventory/{:isbn}/copies", TODO).Methods(http.MethodPost)
	// Remove copy of book
	router.Handle("/api/inventory/{:isbn}/copies/{:id}", TODO).Methods(http.MethodDelete)
	// Check out book
	router.Handle("/api/inventory/{:isbn}/copies/{:id}/checkouts", TODO).Methods(http.MethodPost)
	// Check in book
	router.Handle("/api/inventory/{:isbn}/copies/{:id}/checkins", TODO).Methods(http.MethodPost)

	http.ListenAndServe("/", router)
}

var TODO = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Panic("TODO handler not implemented.")
})
