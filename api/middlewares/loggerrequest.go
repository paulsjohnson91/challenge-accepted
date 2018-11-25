package middlewares

import (
	"log"
	"net/http"

	"github.com/sirupsen/logrus"

	db "../dbs"
	model "../models"
)

var logger *logrus.Logger

func init() {
	log.Printf("[LoggerRequest] loaded!")
	logger = db.Logger()
}

// LoggerRequest middleware for logger all request maded
func LoggerRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		claims, ok := r.Context().Value(model.JwtKey).(model.Claims)
		if !ok {
			claims.UserID = "Unknown"
		}

		logger.WithFields(logrus.Fields{
			"user_id":  claims.UserID,
			"method":   r.Method,
			"endpoint": r.URL.RequestURI(),
		}).Info("LoggerRequest")

		next.ServeHTTP(w, r)

	})
}
