package models

import (
	"gopkg.in/mgo.v2/bson"
)

//Favourites model
type Favourites struct {
	UserID        bson.ObjectId  `json:"id" bson:"_id"`
	ChallengeIDs   []bson.ObjectId  `json:"challengeid" bson:"challengeid"`
}
