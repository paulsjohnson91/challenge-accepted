package middlewares

import (
	"fmt"
	"log"
	"net/http"

	model "../models"
	service "../services"
)

func UserIsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Getting permission")
		claims, ok := r.Context().Value(model.JwtKey).(model.Claims)
		if !ok {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(500)
			fmt.Fprintf(w, `{"message":"Error on decode Context JWT"}`)
			return
		}

		// Check in future
		// mgos, ok := r.Context().Value(model.DbKey).(*mgo.Database)
		// if !ok {
		// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 	w.WriteHeader(500)
		// 	fmt.Fprintf(w, `{"message":"Error on decode Session MongoDB"}`)
		// 	return
		// }

		log.Printf("[UserHavePermission] method=%s EndPoint=%s", r.Method, r.URL.RequestURI())

		if service.UserAdminStatus(claims.UserID) {
			next.ServeHTTP(w, r)
		}

		w.WriteHeader(http.StatusUnauthorized)
		return
	})
}
