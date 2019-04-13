package middlewares

import (
	"github.com/sirupsen/logrus"
	"net/http"

	// "github.com/paulsjohnson91/challenge-accepted/logger"

	model "github.com/paulsjohnson91/challenge-accepted/models"
	"github.com/paulsjohnson91/challenge-accepted/logger"
)

	var log = logger.Logger()

// var logger *logger.Logger

// func init() {
// 	// log.Info("[LoggerRequest] loaded!")
// 	logger = logger.Logger()
// }

// LoggerRequest middleware for logger all request maded
func LoggerRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		claims, ok := r.Context().Value(model.JwtKey).(model.Claims)
		if !ok {
			claims.UserID = "Unknown"
		}

		log.WithFields(logrus.Fields{
			"user_id":  claims.UserID,
			"method":   r.Method,
			"endpoint": r.URL.RequestURI(),
		}).Info("LoggerRequest")

		next.ServeHTTP(w, r)

	})
}
