package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ChallengeItem struct {
	ID    bson.ObjectId   `json:"id" bson:"_id"`
	Index int    		  `json:"index" bson:"index"`
	Item  string 		  `json:"item", bson:"item"`
}

//Project model
type BasicChallenge struct {
	ID             bson.ObjectId   `json:"id" bson:"_id"`
	Name           string          `json:"name" bson:"name"`
	Description    string          `json:"description" bson:"description"`
	Challengeitems []ChallengeItem `json:"challengeitems" bson:"challengeitems"`
	CreatedAt      time.Time       `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at" bson:"updated_at"`
}
