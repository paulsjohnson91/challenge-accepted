package middlewares

import (
	"context"
	"log"
	"net/http"
	"time"

	model "../models"
	mgo "gopkg.in/mgo.v2"
)

const SERVER = "ds055885.mlab.com:55885"

// DBNAME the name of the DB instance
const DBNAME = "gorest"

// DOCNAME the name of the document
const DOCNAME = "albums"

const AuthUserName = "adminuser"

const AuthPassword = "admpass1"

// MongoMiddleware adds mgo MongoDB to context
func MongoMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("MongoDB on request!")

		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    []string{SERVER},
			Timeout:  60 * time.Second,
			Database: DBNAME,
			Username: AuthUserName,
			Password: AuthPassword,
		}

		mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
		if err != nil {
			log.Fatalf("[MongoDB] CreateSession: %s\n", err)
		}
		mongoSession.SetMode(mgo.Monotonic, true)

		rs := mongoSession.Clone()
		defer rs.Close()

		db := rs.DB("gorest")
		ctx := context.WithValue(r.Context(), model.DbKey, db)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
