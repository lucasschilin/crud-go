package entity

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserEntity struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
	Age      int8          `bson:"age"`
}
