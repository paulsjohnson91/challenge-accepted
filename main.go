package main

import (
	"net/http"
	"github.com/joho/godotenv"
	db "github.com/paulsjohnson91/challenge-accepted/dbs"
	route "github.com/paulsjohnson91/challenge-accepted/routes"
	lib "github.com/paulsjohnson91/challenge-accepted/shared"
	"github.com/paulsjohnson91/challenge-accepted/logger"
)

	var log = logger.Logger()

func init() {
	if err := godotenv.Load(); err != nil {
		log.Info("Error loading .env file! Try get a path...")
		if err2 := godotenv.Load(lib.GetPath() + "/.env"); err2 != nil {
			log.Info("Fail...")
			//os.Exit(1)
		}
	}

}

func main() {
	sessions := db.StartDispatch()
	// addr := os.Getenv("API_URL")
	addr := ":3333"

	log.Infof("[Server] Starting server on port %s\n", addr)
	http.ListenAndServe(addr, route.Router(sessions))
}
