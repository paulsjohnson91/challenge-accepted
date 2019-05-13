package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/paulsjohnson91/challenge-accepted/dbs"
	model "github.com/paulsjohnson91/challenge-accepted/models"
	service "github.com/paulsjohnson91/challenge-accepted/services"
	"github.com/paulsjohnson91/challenge-accepted/logger"
)

	var log = logger.Logger()


func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Content-Type", "application/json")
}	

//Home a home API
func Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setupCORS(&w, r)
		w.Write([]byte("Home page"))
	}
}

//Auth get a valid token and expire
func Auth(s *db.Dispatch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setupCORS(&w, r)
		if (*r).Method == "OPTIONS"{
			return
		}
		var user model.User
		decoder := json.NewDecoder(r.Body)
		errDecoder := decoder.Decode(&user)

		if errDecoder != nil {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, `{"message":"Incorrect Decode JSON on body"}`)
			return
		}

		t, err := service.GenerateToken(s, user)
		if err != nil {
			log.Info("Error : %q", err)
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, `{"message": %q}`, err)
			return
		}
		log.Info(t)
		//write json
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"token":%q, "expire":%s, "firstname":"%s", "lastname":"%s", "admin":"%t", "userid":"%s"}`, t.Token, t.Expire, t.FirstName, t.LastName, t.Admin, t.UserID)
	}
}
