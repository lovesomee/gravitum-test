package api

import (
	"encoding/json"
	"gravitum-test/models"
	"log"
	"net/http"
)

func UpdateUser(userService UserService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var users models.Users
		if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
			log.Println(err)
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if err := userService.UpdateUser(users); err != nil {
			log.Println(err)
			http.Error(w, "failed to update user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
