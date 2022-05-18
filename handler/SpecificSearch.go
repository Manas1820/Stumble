package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"stumble/models"
)

// Use use the like feature of the Postgres

func (h handler) GetSearchBasedUser(w http.ResponseWriter, r *http.Request) {

	// Read dynamic search parameter
	search_param := r.Header.Get("search")

	var user []models.User

	search_string := "%" + search_param + "%"

	if result := h.DB.Where("name LIKE ?", search_string).Find(&user); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
