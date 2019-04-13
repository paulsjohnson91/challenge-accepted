package middlewares

import (
	"fmt"
	"net/http"

	model "github.com/paulsjohnson91/challenge-accepted/models"
	service "github.com/paulsjohnson91/challenge-accepted/services"
)

func UserIsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[UserIsAdmin] Getting permission")
		claims, ok := r.Context().Value(model.JwtKey).(model.Claims)
		if !ok {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(500)
			fmt.Fprintf(w, `{"message":"Error on decode Context JWT"}`)
			return
		}

		log.Info("[UserIsAdmin] method=%s EndPoint=%s", r.Method, r.URL.RequestURI())

		if service.UserAdminStatus(claims.UserID) {
			next.ServeHTTP(w, r)
		} else {
			log.Info("[UserIsAdmin] Unauthorized access to method=%s EndPoint=%s by user=%s", r.Method, r.URL.RequestURI(), claims.UserID)
		}

		w.WriteHeader(http.StatusUnauthorized)
		return
	})
}
