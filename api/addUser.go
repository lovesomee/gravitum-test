package api

import (
	"encoding/json"
	"gravitum-test/models"
	"log"
	"net/http"
)

func AddUser(userService UserService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var users models.Users
		if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
		}

		if err := userService.AddUser(users); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
		}

		w.WriteHeader(http.StatusCreated)
	}
}
