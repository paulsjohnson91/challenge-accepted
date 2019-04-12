package main

import (
	"fmt"
	"net/http"

	"os"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	db "github.com/paulsjohnson91/challenge-accepted/dbs"
	route "github.com/paulsjohnson91/challenge-accepted/routes"
	lib "github.com/paulsjohnson91/challenge-accepted/shared"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "15:04:05",
		FullTimestamp:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			repopath := fmt.Sprintf("%s/src/github.com/paulsjohnson91", os.Getenv("GOPATH"))
			filename := strings.Replace(f.File, repopath, "", -1)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

	log.Info("Init")
	// load config file
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

	log.Info("[Server] Path: %s", lib.GetPath())
	fmt.Printf("[Server] Starting server on %v\n", addr)
	http.ListenAndServe(addr, route.Router(sessions))
}
