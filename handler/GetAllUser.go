package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"stumble/models"
)

func (h handler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	if result := h.DB.Find(&users); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
