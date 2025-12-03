package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"gravitum-test/config"
)

func NewServer(cfg config.Settings, logger *zap.Logger, users UserService) *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/ping", Ping(logger)).Methods(http.MethodGet)
	router.HandleFunc("/users", AddUser(logger, users)).Methods(http.MethodPost)
	router.HandleFunc("/users", UpdateUser(logger, users)).Methods(http.MethodPut)
	router.HandleFunc("/users", GetUser(logger, users)).Methods(http.MethodGet)

	return &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", cfg.Port),
	}
}
