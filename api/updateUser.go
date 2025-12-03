package api

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"gravitum-test/models"
)

func UpdateUser(logger *zap.Logger, userService UserService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var users models.Users
		if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
			logger.Warn("invalid request body", zap.Error(err))
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if err := userService.UpdateUser(ctx, users); err != nil {
			logger.Error("failed to update user", zap.Error(err), zap.Int("user_id", users.Id))
			http.Error(w, "failed to update user", http.StatusInternalServerError)
			return
		}

		logger.Info("user updated", zap.Int("user_id", users.Id))
		w.WriteHeader(http.StatusOK)
	}
}
