package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"gravitum-test/config"
	"net/http"
)

func NewServer(cfg config.Settings, users UserService) *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/ping", Ping()).Methods(http.MethodGet)
	router.HandleFunc("/users", AddUser(users)).Methods(http.MethodPost)
	router.HandleFunc("/users", UpdateUser(users)).Methods(http.MethodPut)
	router.HandleFunc("/users", GetUser(users)).Methods(http.MethodGet)

	return &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", cfg.Port),
	}
}
