package models

import (
	"gopkg.in/mgo.v2/bson"
)

//Favourites model
type Favourites struct {
	ID            bson.ObjectId  `json:"id" bson:"_id"`
	ChallengeID   bson.ObjectId  `json:"challengeid" bson:"challengeid"`
	UserID        bson.ObjectId  `json:"userid" bson:"userid"`
}
