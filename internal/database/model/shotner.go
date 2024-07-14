package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Shotner struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Short  string             `json:"short,omitempty"`
	Origin string             `json:"origin,omitempty"`
}
