package middlewares

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/pressly/chi"

	model "github.com/paulsjohnson91/challenge-accepted/models"
	service "github.com/paulsjohnson91/challenge-accepted/services"
)

// UserValidOnProject middleware for validate permission of user
func UserValidOnProject(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("permissions being checked")
		slug := chi.URLParam(r, "slug")

		if slug == "" {
			log.Info("[UserValidOnProject] Slug empty!")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims, ok := r.Context().Value(model.JwtKey).(model.Claims)
		if !ok {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(500)
			fmt.Fprintf(w, `{"message":"Error on decode Context JWT - middleware UserValidOnProject "}`)
			return
		}

		log.Info("[UserValidOnProject] method=%s EndPoint=%s SLUG=%s", r.Method, r.URL.RequestURI(), slug)

		project, err := service.UserIsValidOnProject(slug, claims.UserID)
		if err != nil {
			log.Info("[UserValidOnProject] Err: %s", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), model.ProjKey, project.ID.Hex())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
