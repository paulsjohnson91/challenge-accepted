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
	model "github.com/paulsjohnson91/challenge-accepted/models"
	lib "github.com/paulsjohnson91/challenge-accepted/shared"
)

//GetUser get a user by Id
func GetUser(s *db.Dispatch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ss := s.MongoDB.Copy()
		defer ss.Close()

		// Grab id
		id := chi.URLParam(r, "id")

		// Verify id is ObjectId, otherwise bail
		if !bson.IsObjectIdHex(id) {
			log.Println("ID is not BSON ID")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		oid := bson.ObjectIdHex(id)
		u := model.User{}
		if err := ss.DB("gorest").C("users").FindId(oid).One(&u); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		uj, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", uj)
	}
}

func GetUsers(s *db.Dispatch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setupCORS(&w, r)
		// if (*r).Method == "OPTIONS"{
		// 	return
		// }
		ss := s.MongoDB.Copy()
		defer ss.Close()

		u := []model.User{}
		if err := ss.DB("gorest").C("users").Find(nil).All(&u); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		uj, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", uj)
	}
}

func doesUserExist(s *db.Dispatch, email string) bool {
	ss := s.MongoDB.Copy()
	u := model.User{}
	if err := ss.DB("gorest").C("users").Find(bson.M{"email": email}).One(&u); err != nil {
		return false
	}
	return true
}

//CreateUser create a new user
func CreateUser(s *db.Dispatch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ss := s.MongoDB.Copy()
		defer ss.Close()

		u := model.User{}
		json.NewDecoder(r.Body).Decode(&u)
        if doesUserExist(s,u.Email){
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, "User already exists:%s",u.Email)
			return
		}
		// Add an Id
		u.ID = bson.NewObjectId()
		u.CreatedAt = time.Now()
		u.UpdatedAt = time.Now()

		if passwd, err := lib.Encrypt(u.Password); err == nil {
			u.Password = passwd
		}
		u.Admin = false
		names, err := ss.DB("gorest").CollectionNames()
		if err != nil {
			// Handle error
			log.Info("Failed to get coll names: %v", err)
			return
		}

		isUsersTable := false
		isUsers := false
		for _, name := range names {
			if name == "users" {
				isUsersTable = true
				count, err := ss.DB("gorest").C("users").Count()
				if err != nil {
					// Handle error
					log.Info("Count on users table failed", err)
					return
				}

				if count == 0 {
					isUsers = true
				}
				break
			}
		}
		if isUsers == true || isUsersTable == false {
			log.Info("There are no users, adding user as admin")
			u.Admin = true
		}

		ss.DB("gorest").C("users").Insert(u)
		uj, _ := json.Marshal(u)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s", uj)
	}
}

//CreateAdmin create a new user
func CreateAdminUser(s *db.Dispatch) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ss := s.MongoDB.Copy()
		defer ss.Close()

		u := model.User{}
		json.NewDecoder(r.Body).Decode(&u)

		// Add an Id
		u.ID = bson.NewObjectId()
		u.CreatedAt = time.Now()
		u.UpdatedAt = time.Now()

		if passwd, err := lib.Encrypt(u.Password); err == nil {
			u.Password = passwd
		}
		u.Admin = true

		ss.DB("gorest").C("users").Insert(u)
		uj, _ := json.Marshal(u)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s", uj)
	}
}

// DeleteUser remove user from database
func DeleteUser(s *db.Dispatch) http.HandlerFunc {
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

		c := ss.DB("gorest").C("users")

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

// UpdateUser update user
func UpdateUser(s *db.Dispatch) http.HandlerFunc {
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

		// Stub an user to be populated from the body
		u := model.User{}
		json.NewDecoder(r.Body).Decode(&u)
		u.UpdatedAt = time.Now()

		c := ss.DB("gorest").C("users")

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
