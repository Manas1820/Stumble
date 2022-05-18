package router

import (
	db "stumble/database"
	handlers "stumble/handler"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go

func Router() *mux.Router {

	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	// Get a specific user by its id
	router.HandleFunc("/api/user/{id}", h.GetUser).Methods("GET", "OPTIONS")

	// router.HandleFunc("/api/user", h.GetAllUsers).Methods("GET", "OPTIONS")

	// Handle the search
	router.HandleFunc("/api/user/", h.GetSearchBasedUser).Methods("GET", "OPTIONS")

	// Handle the distance search
	router.HandleFunc("/api/bonding", h.GetDistanceBasedUser).Methods("GET", "OPTIONS")

	return router
}
