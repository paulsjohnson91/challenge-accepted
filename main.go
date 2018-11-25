package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"

	"log"

	route "./routes"
	lib "./shared"
	db "github.com/paulsjohnson91/challenge-accepted/dbs"
)

func init() {
	log.Println("Init")
	// load config file
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file! Try get a path...")
		if err2 := godotenv.Load(lib.GetPath() + "/.env"); err2 != nil {
			log.Printf("Fail...")
			//os.Exit(1)
		}
	}
}

func main() {
	fmt.Println("here")
	log.Println("Starting")
	sessions := db.StartDispatch()
	// addr := os.Getenv("API_URL")
	addr := ":3333"

	log.Printf("[Server] Path: %s", lib.GetPath())
	fmt.Printf("[Server] Starting server on %v\n", addr)
	http.ListenAndServe(addr, route.Router(sessions))
}
