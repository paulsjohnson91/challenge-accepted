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

//Get Progress of Subscription
func GetProgress(s *db.Dispatch) http.HandlerFunc {
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
		u := model.Subscription{}
		if err := ss.DB("gorest").C("subscriptions").Find(bson.M{"userid": bson.ObjectIdHex(claims.UserID), "challengeid": oid}).One(&u); err != nil {
			progress := model.Progress{}
			progress.Progress = 0.0
			progress.Active = false
			uj, _ := json.Marshal(progress)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s", uj)
			return
		}
		items := 0.0
		itemsComplete := 0.0
        for _, it := range u.ItemsProgress {
			if it.Complete == true {
				itemsComplete = itemsComplete + 1
			}
			items = items + 1
		}

		progress := model.Progress{}
		progress.Active = true
		progress.Progress = itemsComplete * 100 / items

		uj, _ := json.Marshal(progress)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", uj)
	}
}

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

		items := 0.0
		itemsComplete := 0.0
        for _, it := range u.ItemsProgress {
			if it.Complete == true {
				itemsComplete = itemsComplete + 1
			}
			items = items + 1
		}

		u.Progress = itemsComplete * 100 / items

		uj, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", uj)
	}
}

//GetChallenge get a challenge by Id
func GetSubscriptionByCID(s *db.Dispatch) http.HandlerFunc {
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
		u := model.Subscription{}
		if err := ss.DB("gorest").C("subscriptions").Find(bson.M{"userid": bson.ObjectIdHex(claims.UserID), "challengeid": oid}).One(&u); err != nil {
			log.Info(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		items := 0.0
		itemsComplete := 0.0
        for _, it := range u.ItemsProgress {
			if it.Complete == true {
				itemsComplete = itemsComplete + 1
			}
			items = items + 1
		}

		u.Progress = itemsComplete * 100 / items

		uj, _ := json.Marshal(u)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", uj)
	}
}

//GetSubscriptions get all subscritions for user
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

//GetCompletedSubscriptions get all subscritions for user
func GetCompletedSubscriptions(s *db.Dispatch) http.HandlerFunc {
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
		if err := ss.DB("gorest").C("subscriptions").Find(bson.M{"userid": bson.ObjectIdHex(claims.UserID),"iscomplete": true}).All(&u); err != nil {
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

		isSubscription := model.Subscription{}
		if err := ss.DB("gorest").C("subscriptions").Find(bson.M{"userid": bson.ObjectIdHex(claims.UserID), "challengeid": oid}).One(&isSubscription); err != nil {
			log.Info("No subscription found for challenge " + id + " creating new one")
		} else {
			if(isSubscription.IsComplete == true){
				isSubscription.IsComplete = false
				for i, _ := range isSubscription.ItemsProgress {
					isSubscription.ItemsProgress[i].Complete = false
				}
				c := ss.DB("gorest").C("subscriptions")

				if err := c.Update(bson.M{"userid": bson.ObjectIdHex(claims.UserID), "challengeid": oid}, &isSubscription); err != nil {
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
				}
			uj, _ := json.Marshal(isSubscription)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "%s", uj)
			return
		}

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
		u.Progress = 0

		ss.DB("gorest").C("subscriptions").Insert(u)
		uj, _ := json.Marshal(u)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s", uj)
	}
}

//CreateChallenge create challenge
func UpdateItem(s *db.Dispatch) http.HandlerFunc {
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

		subitem := model.SubscriptionItem{}
		json.NewDecoder(r.Body).Decode(&subitem)



		// Grab id
		id := chi.URLParam(r, "id")
		itemid := chi.URLParam(r, "itemid")


		// Verify id is ObjectId, otherwise bail
		if !bson.IsObjectIdHex(id) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		oid := bson.ObjectIdHex(id)
		iid := bson.ObjectIdHex(itemid)
		u := model.Subscription{}
		if err := ss.DB("gorest").C("subscriptions").FindId(oid).One(&u); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if u.IsComplete == true {
			w.WriteHeader(http.StatusConflict)
			return
		}

		if bson.ObjectIdHex(claims.UserID) != u.UserID {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		for i, element := range u.ItemsProgress {
			if element.ChallengeItemID == iid {
				if(subitem.Complete == "true"){
					u.ItemsProgress[i].Complete = true
					u.ItemsProgress[i].CompletedAt = time.Now()
					
				} else{
					u.ItemsProgress[i].Complete = false
				}
				u.UpdatedAt = time.Now()
				
			}
		}
		items := 0.0
		itemsComplete := 0.0
        for _, it := range u.ItemsProgress {
			if it.Complete == true {
				itemsComplete = itemsComplete + 1
			}
			items = items + 1
		}

		u.Progress = itemsComplete * 100 / items
		if u.Progress == 100 {
			u.IsComplete = true
			u.TimesCompleted = u.TimesCompleted + 1
			u.LastCompleted = time.Now()
		}
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

		if u.IsComplete == true {
			user := basemodel.User{}
			if err := ss.DB("gorest").C("users").FindId(bson.ObjectIdHex(claims.UserID)).One(&user); err != nil {
				switch err {
				default:
					msg := []byte(`{"message":"Could not update user progress"}`)
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintf(w, "%s", msg)
	
				case mgo.ErrNotFound:
					msg := []byte(`{"message":"Could not update user progress"}`)
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "%s", msg)
				}
				return
			} else {
				user.Completed = user.Completed + 1
				if err := ss.DB("gorest").C("users").Update(bson.M{"_id": bson.ObjectIdHex(claims.UserID)}, &user); err != nil {
					switch err {
					default:
						msg := []byte(`{"message":"Could not update user progress"}`)
						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(http.StatusInternalServerError)
						fmt.Fprintf(w, "%s", msg)
		
					case mgo.ErrNotFound:
						msg := []byte(`{"message":"Could not update user progress"}`)
						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(http.StatusNotFound)
						fmt.Fprintf(w, "%s", msg)
					}
					return
				}
			}
		
		}

		uj, _ := json.Marshal(u)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
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
