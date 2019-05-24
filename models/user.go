package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User model
type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Email     string        `json:"email" bson:"email"`
	Password  string        `json:"password" bson:"password"`
	FirstName string        `json:"firstname" bson:"firstname"`
	LastName  string        `json:"lastname" bson:"lastname"`
	Admin     bool          `json:"admin" bson:"admin"`
	Completed int           `json:"completed" bson:"completed"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}
