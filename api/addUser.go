package api

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"gravitum-test/models"
)

func AddUser(logger *zap.Logger, userService UserService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var users models.Users
		if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
			logger.Warn("invalid request body", zap.Error(err))
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if err := userService.AddUser(ctx, users); err != nil {
			logger.Error("failed to add user", zap.Error(err))
			http.Error(w, "failed to add user", http.StatusInternalServerError)
			return
		}

		logger.Info("user added", zap.String("first_name", users.FirstName), zap.String("last_name", users.LastName))
		w.WriteHeader(http.StatusCreated)
	}
}
