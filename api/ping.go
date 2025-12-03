package api

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func Ping(logger *zap.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Debug("ping request", zap.String("remote_addr", r.RemoteAddr))
		fmt.Fprintf(w, "pong")
	}
}
