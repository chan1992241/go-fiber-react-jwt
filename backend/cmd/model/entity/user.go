package entity

import (
	mongodb "chan1992241/backend/cmd/model/bean"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	Username     string             `bson:"username"`
	Password     string             `bson:"password"`
	RefreshToken string             `bson:"refreshToken"`
	Role         string             `bson:"role"`
}

var UserCollection = mongodb.MongoDatabase.Collection("users")
