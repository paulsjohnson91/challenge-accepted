package routes

import (
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/rs/cors"

	controller "github.com/paulsjohnson91/challenge-accepted/controllers"
	db "github.com/paulsjohnson91/challenge-accepted/dbs"
	mid "github.com/paulsjohnson91/challenge-accepted/middlewares"
)

//LoggedIn Routes
func LoggedIn(s *db.Dispatch, cors *cors.Cors) func(r chi.Router) {
	return func(r chi.Router) {
		r.Use(middleware.DefaultCompress)
		r.Use(middleware.RequestID)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Use(cors.Handler)
		r.Use(mid.LoggerRequest)
		//Chain of validation user
		r.Use(mid.TokenAuthentication) //If token ok

		r.Get("/challenge/{id}", controller.GetChallenge(s))
		r.Get("/challenge", controller.GetChallenges(s))

		r.Get("/subscription/{id}", controller.GetSubscription(s))
		r.Get("/subscription", controller.GetSubscriptions(s))
		r.Post("/subscription/{id}", controller.CreateSubscription(s))
		r.Put("/subscription/{id}", controller.UpdateSubscription(s))
		r.Delete("/subscription/{id}", controller.DeleteSubscription(s))
		r.Post("/subscription/{id}/{itemid}", controller.UpdateItem(s))
		r.Get("/progress/{id}", controller.GetProgress(s))


		r.Get("/favourites", controller.GetFavourites(s))
		r.Post("/favourites/{id}", controller.CreateFavourite(s))
		r.Delete("/favourites/{id}", controller.DeleteFavourite(s))

	}
}
