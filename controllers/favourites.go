package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/pressly/chi"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	db "github.com/paulsjohnson91/challenge-accepted/dbs"
	basemodel "github.com/paulsjohnson91/challenge-accepted/models"
	model "github.com/paulsjohnson91/challenge-accepted/models/challenges"
)

//Get all favourites for user
func GetFavourites(s *db.Dispatch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ss := s.MongoDB.Copy()
		defer ss.Close()

		claims, ok := r.Context().Value(basemodel.JwtKey).(basemodel.Claims)
		if !ok {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(500)
			fmt.Fprintf(w, `{"message":"Error on decode Context JWT"}`)
			return
		}

		u := []model.Favourites{}
		ss.DB("gorest").C("favourites").Find(bson.M{"userid": bson.ObjectIdHex(claims.UserID)}).All(&u)

		uj, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", uj)
	}
}

//GetSubscritions get all subscritions for all user
func GetAllFavourites(s *db.Dispatch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ss := s.MongoDB.Copy()
		defer ss.Close()

		u := []model.Favourites{}
		ss.DB("gorest").C("favourites").Find(nil).All(&u)

		uj, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", uj)
	}
}

//CreateChallenge create challenge
func CreateFavourite(s *db.Dispatch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ss := s.MongoDB.Copy()
		defer ss.Close()

		claims, ok := r.Context().Value(basemodel.JwtKey).(basemodel.Claims)
		if !ok {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(500)
			fmt.Fprintf(w, `{"message":"Error on decode Context JWT"}`)
			return
		}

		id := chi.URLParam(r, "id")

		// Verify id is ObjectId, otherwise bail
		if !bson.IsObjectIdHex(id) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		oid := bson.ObjectIdHex(id)


		u := model.Favourites{}

		// uu := []model.Favourites{}
		// ss.DB("gorest").C("favourites").FindId(oid).All(&uu)

		count, _ := ss.DB("gorest").C("favourites").Find(bson.M{"userid": bson.ObjectIdHex(claims.UserID), "challengeid": oid}).Count()

		if(count > 0){
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, "%s", "Favourite already exists")
			return
		}

		// Add an Id
		u.ID = bson.NewObjectId()
		u.ChallengeID = oid
		u.UserID = bson.ObjectIdHex(claims.UserID)
		ss.DB("gorest").C("favourites").Insert(u)
		uj, _ := json.Marshal(u)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s", uj)
	}
}

func DeleteFavourite(s *db.Dispatch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ss := s.MongoDB.Copy()
		defer ss.Close()

		id := chi.URLParam(r, "id")

		if !bson.IsObjectIdHex(id) {
			msg := []byte(`{"message":"ObjectId invalid"}`)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "%s", msg)
			return
		}

		c := ss.DB("gorest").C("favourites")

		if err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)}); err != nil {
			switch err {
			default:
				msg := []byte(`{"message":"ObjectId invalid"}`)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "%s", msg)

			case mgo.ErrNotFound:
				msg := []byte(`{"message":"ObjectId not found"}`)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "%s", msg)
			}
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
