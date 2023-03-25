package controller

import (
	"chan1992241/backend/cmd/model/entity"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	var hashedPassword []byte
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	hashedPassword, _ = bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	var result, err = entity.UserCollection.InsertOne(c.Context(), entity.User{
		Username: data["username"],
		Password: string(hashedPassword),
	})
	if err != nil {
		return err
	}
	fmt.Println(result.InsertedID)
	return c.JSON(hashedPassword)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user entity.User
	filter := bson.D{{Key: "username", Value: data["username"]}}
	err := entity.UserCollection.FindOne(c.Context(), filter).Decode(&user)
	if err != nil {
		return err
	}
	return c.JSON(user)
}
