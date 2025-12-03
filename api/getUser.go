package api

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func GetUser(logger *zap.Logger, userService UserService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		users, err := userService.GetUser(ctx)
		if err != nil {
			logger.Error("failed to get users", zap.Error(err))
			http.Error(w, "failed to get users", http.StatusInternalServerError)
			return
		}

		resp, err := json.Marshal(users)
		if err != nil {
			logger.Error("failed to marshal users", zap.Error(err))
			http.Error(w, "failed to marshal users", http.StatusInternalServerError)
			return
		}

		logger.Info("users fetched", zap.Int("count", len(users)))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}
