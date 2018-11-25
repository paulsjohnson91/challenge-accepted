package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ItemProgress struct {
	Index       int       `json:"index" bson:"index"`
	Complete    bool      `json:"complete", bson:"complete"`
	CompletedAt time.Time `json:"completed_at,omitempty" bson:"completed_at"`
}

//Project model
type Subscription struct {
	ID            bson.ObjectId  `json:"id" bson:"_id"`
	ChallengeID   bson.ObjectId  `json:"challengeid" bson:"challengeid"`
	UserID        bson.ObjectId  `json:"userid" bson:"userid"`
	ItemsProgress []ItemProgress `json:"itemsprogress" bson:"itemsprogress"`
	CreatedAt     time.Time      `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" bson:"updated_at"`
	IsComplete    bool           `json:"iscomplete" bson:"iscomplete"`
}
