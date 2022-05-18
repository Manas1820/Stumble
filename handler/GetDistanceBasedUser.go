package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"stumble/models"
)

// Use use the like feature of the Postgres

func (h handler) GetDistanceBasedUser(w http.ResponseWriter, r *http.Request) {

	// Read dynamic search parameter
	given_user, _ := strconv.Atoi(r.Header.Get("user_id"))
	location, _ := strconv.ParseFloat(r.Header.Get("location"), 64)

	// check for valid user
	var given_user_deatils models.User

	if result := h.DB.First(&given_user_deatils, given_user); result.Error != nil {
		fmt.Println(result.Error)
	}

	//check for valid location

	lower_location := given_user_deatils.Location - location
	upper_location := given_user_deatils.Location + location

	var user []models.User

	// remove the present user from the search params
	if result := h.DB.Where("location BETWEEN ? AND ? AND NOT id=?", lower_location, upper_location, given_user).Find(&user); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
