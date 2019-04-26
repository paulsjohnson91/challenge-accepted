package routes

import (
	"net/http"

	"github.com/pressly/chi"
	"github.com/rs/cors"
	"github.com/paulsjohnson91/challenge-accepted/logger"
	db "github.com/paulsjohnson91/challenge-accepted/dbs"
)
var log = logger.Logger()
func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Content-Type", "application/json")
}
//All options requests
func AllOptions(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  if r.Method == "OPTIONS" {
		setupCORS(&w, r)
		w.WriteHeader(http.StatusOK)
		return // we return because likely we want to finish the request here.
	  }
	  next.ServeHTTP(w, r) // keep processing for non-OPTIONS methods
	})
  }

//Router main rules of routes
func Router(s *db.Dispatch) http.Handler {
	r := chi.NewRouter()

	//CORS setup
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		// AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowedHeaders:   []string{"*"},
        OptionsPassthrough: true,
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(AllOptions)

	// Protected routes
	r.Group(Protected(s, cors))
	// Public routes
	r.Group(Public(s, cors))

	r.Group(LoggedIn(s, cors))

	return r
}
