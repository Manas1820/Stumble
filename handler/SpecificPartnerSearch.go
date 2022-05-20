package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"stumble/models"

	"github.com/gorilla/mux"
)

func (h handler) SpecificPartnerSearch(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var likes []models.Likes

	if result := h.DB.Find(&likes, "who_likes = ?", id); result.Error != nil {
		fmt.Println(result.Error)
	}

	var liked_user_details []models.User

	fmt.Println(len(likes))
	for i := 0; i < len(likes); i++ {

		var likes_user models.Likes
		if result := h.DB.Find(&likes_user, "who_likes = ? AND who_is_liked = ?", id, likes[i].WhoIsLiked); result.Error != nil {

		}

		var temp []models.User
		if likes_user.ID > 0 {
			if result := h.DB.First(&temp, likes[i].WhoIsLiked); result.Error != nil {
				fmt.Println(result.Error)
			}

			liked_user_details = append(temp)
		}

	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(liked_user_details)
}
