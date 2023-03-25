package entity

import (
	mongodb "chan1992241/backend/cmd/model/bean"
)

type User struct {
	Username     string `bson:"username"`
	Password     string `bson:"password"`
	RefreshToken string `bson:"refreshToken"`
}

var UserCollection = mongodb.MongoDatabase.Collection("users")
