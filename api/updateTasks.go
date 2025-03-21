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
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
		}

		if err := userService.UpdateUser(users); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
		}

		w.WriteHeader(http.StatusOK)
	}
}
