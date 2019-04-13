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

//Home a home API
func Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Home page"))
	}
}

//Auth get a valid token and expire
func Auth(s *db.Dispatch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		decoder := json.NewDecoder(r.Body)
		errDecoder := decoder.Decode(&user)

		if errDecoder != nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, `{"message":"Incorrect Decode JSON on body"}`)
			return
		}

		t, err := service.GenerateToken(s, user)
		if err != nil {
			log.Info("Error : %q", err)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, `{"message": %q}`, err)
			return
		}

		//write json
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"token":%q, "expire":%s}`, t.Token, t.Expire)
	}
}
