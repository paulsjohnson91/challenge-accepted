package routes

import (
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/rs/cors"

	controller "github.com/paulsjohnson91/challenge-accepted/controllers"
	db "github.com/paulsjohnson91/challenge-accepted/dbs"
	mid "github.com/paulsjohnson91/challenge-accepted/middlewares"
)

//Protected Routes
func Protected(s *db.Dispatch, cors *cors.Cors) func(r chi.Router) {
	return func(r chi.Router) {
		r.Use(middleware.DefaultCompress)
		r.Use(middleware.RequestID)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Use(cors.Handler)
		//Chain of validation user
		r.Use(mid.TokenAuthentication) //If token ok
		//r.Use(mid.UserValidOnProject)  //if user belong to project ok
		//r.Use(mid.UserHavePermission) //if user has permisson on endpoint ok
		r.Use(mid.LoggerRequest) //log any request ok
		r.Use(mid.UserIsAdmin)
		//endpoint protected
		// r.Get("/admin/:slug", controller.Admin())

		//CRUD User
		r.Get("/user/{id}", controller.GetUser(s))
		r.Put("/user/{id}", controller.UpdateUser(s))
		r.Delete("/user/{id}", controller.DeleteUser(s))

		r.Get("/users", controller.GetUsers(s))

		//CRUD Admin User
		r.Post("/admin", controller.CreateAdminUser(s))

		//Subscriptions
		r.Get("/allsubscription", controller.GetAllSubscriptions(s))

		//CRUD Permission
		// r.Post("/permission", controller.CreatePermission(s))
		// r.Get("/permission/:id", controller.GetPermission(s))
		// r.Put("/permission/:id", controller.UpdatePermission(s))
		// r.Delete("/permission/:id", controller.DeletePermission(s))
		//
		// //CRUD Project
		// r.Post("/project", controller.CreateProject(s))
		// r.Get("/project/:id", controller.GetProject(s))
		// r.Put("/project/:id", controller.UpdateProject(s))
		// r.Delete("/project/:id", controller.DeleteProject(s))
		//
		r.Post("/challenge", controller.CreateChallenge(s))
		r.Delete("/challenge/{id}", controller.DeleteChallenge(s))
		r.Put("/challenge/{id}", controller.UpdateChallenge(s))

	}
}
