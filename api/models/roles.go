package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Permission model
type Role struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	EndPoint  string        `json:"endpoint" bson:"endpoint"`
	Owner     bson.ObjectId `json:"owner" bson:"owner"`
	Project   bson.ObjectId `json:"project" bson:"project"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}
