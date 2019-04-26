package routes

import (
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/rs/cors"
	controller "github.com/paulsjohnson91/challenge-accepted/controllers"
	db "github.com/paulsjohnson91/challenge-accepted/dbs"
	mid "github.com/paulsjohnson91/challenge-accepted/middlewares"
)


//Public Routes
func Public(s *db.Dispatch, cors *cors.Cors) func(r chi.Router) {
	return func(r chi.Router) {	


		r.Use(middleware.DefaultCompress)
		r.Use(middleware.RequestID)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Use(cors.Handler)
		r.Use(mid.LoggerRequest)

		// home
		r.Get("/", controller.Home())

		// Authenticate user
		r.Post("/auth", controller.Auth(s))
		r.Options("/auth", controller.Auth(s))

		//CRUD User
		r.Post("/users/register", controller.CreateUser(s))
	}
}
