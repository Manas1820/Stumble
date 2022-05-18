package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"stumble/models"

	"github.com/gorilla/mux"
)

func (h handler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find user by User ID
	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
