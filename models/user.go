package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Firstname string        `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string        `json:"lastname,omitempty" bson:"lastname,omitempty"`
}
