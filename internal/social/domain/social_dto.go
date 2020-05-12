package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	//StatusActive user create has a default status of active
	StatusActive = "Active"
)

//User ...
type User struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Gender string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Age    int                `json:"age,omitempty" bson:"age,omitempty"`
	State  string             `json:"state,omitempty" bson:"state,omitempty"`
}

//Friendship ..
type Friendship struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	From        string             `json:"from,omitempty" bson:"from,omitempty"`
	To          string             `json:"to,omitempty" bson:"to,omitempty"`
	ConnectDate string             `json:"connect_day,omitempty" bson:"connect_day,omitempty"`
}
