package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/pressly/chi"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	db "github.com/paulsjohnson91/challenge-accepted/dbs"
	basemodel "github.com/paulsjohnson91/challenge-accepted/models"
	model "github.com/paulsjohnson91/challenge-accepted/models/challenges"
)

//GetChallenge get a challenge by Id
func GetSubscription(s *db.Dispatch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ss := s.MongoDB.Copy()
		defer ss.Close()

		// Grab id
		id := chi.URLParam(r, "id")

		// Verify id is ObjectId, otherwise bail
		if !bson.IsObjectIdHex(id) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		oid := bson.ObjectIdHex(id)
		u := model.Subscription{}
		if err := ss.DB("gorest").C("subscriptions").FindId(oid).One(&u); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		uj, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", uj)
	}
}

//GetSubscritions get all subscritions for user
func GetSubscriptions(s *db.Dispatch) http.HandlerFunc {
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

		u := []model.Subscription{}
		log.Info("[getSubscriptions] search by user " + claims.UserID)
		if err := ss.DB("gorest").C("subscriptions").Find(bson.M{"userid": bson.ObjectIdHex(claims.UserID)}).All(&u); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		uj, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", uj)
	}
}

//GetSubscritions get all subscritions for all user
func GetAllSubscriptions(s *db.Dispatch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ss := s.MongoDB.Copy()
		defer ss.Close()

		u := []model.Subscription{}
		if err := ss.DB("gorest").C("subscriptions").Find(nil).All(&u); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		uj, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", uj)
	}
}

//CreateChallenge create challenge
func CreateSubscription(s *db.Dispatch) http.HandlerFunc {
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

		// Grab id
		id := chi.URLParam(r, "id")

		// Verify id is ObjectId, otherwise bail
		if !bson.IsObjectIdHex(id) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		oid := bson.ObjectIdHex(id)
		c := model.BasicChallenge{}
		if err := ss.DB("gorest").C("challenges").FindId(oid).One(&c); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		u := model.Subscription{}
		itemProgress := []model.ItemProgress{}
		// Add an Id
		u.ID = bson.NewObjectId()
		u.CreatedAt = time.Now()
		u.UpdatedAt = time.Now()
		u.UserID = bson.ObjectIdHex(claims.UserID)
		u.ChallengeID = c.ID
		u.IsComplete = false
		for _, element := range c.Challengeitems {
			itemProgress = append(itemProgress, CreateItemProgress(element))
		}
		u.ItemsProgress = itemProgress

		ss.DB("gorest").C("subscriptions").Insert(u)
		uj, _ := json.Marshal(u)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s", uj)
	}
}
//Maps ChallengeItem to ItemProgress
func CreateItemProgress(element model.ChallengeItem) model.ItemProgress {
	itemProgressList := model.ItemProgress{}
	log.Info("[CreateItemProgress] creating item progress for %s" ,element.ID)
	itemProgressList.ChallengeItemID = element.ID
	itemProgressList.Complete = false
	return itemProgressList
}

func DeleteSubscription(s *db.Dispatch) http.HandlerFunc {
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

		c := ss.DB("gorest").C("subscriptions")

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

// Updatechallenge update challenge
func UpdateSubscription(s *db.Dispatch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ss := s.MongoDB.Copy()
		defer ss.Close()

		id := chi.URLParam(r, "id")

		// Verify id is ObjectId, otherwise bail
		if !bson.IsObjectIdHex(id) {
			msg := []byte(`{"message":"ObjectId invalid"}`)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "%s", msg)
			return
		}

		// Stub an challenge to be populated from the body
		u := model.Subscription{}
		json.NewDecoder(r.Body).Decode(&u)
		u.UpdatedAt = time.Now()

		c := ss.DB("gorest").C("subscriptions")

		if err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &u); err != nil {
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

		w.WriteHeader(http.StatusOK)
	}
}
