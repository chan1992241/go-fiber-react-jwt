package controller

import (
	"chan1992241/backend/cmd/model/entity"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	var hashedPassword []byte
	if err := c.BodyParser(&data); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	hashedPassword, _ = bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	var result, err = entity.UserCollection.InsertOne(c.Context(), entity.User{
		ID:       primitive.NewObjectID(),
		Username: data["username"],
		Password: string(hashedPassword),
		Role:     data["role"],
	})
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(result)
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
		c.Status(fiber.StatusNotFound)
		return c.SendString("User not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.SendString("Incorrect password")
	}
	claims := jwt.MapClaims{
		"userId": user.ID,
		//set one hour expiration
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.Expires = time.Now().Add(1 * time.Hour)
	c.Cookie(cookie)
	// Create refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	signedRefreshToken, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	user.RefreshToken = string(signedRefreshToken)
	_, err = entity.UserCollection.UpdateOne(c.Context(), filter, bson.D{{Key: "$set", Value: user}})
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"token": t})
}

func RefreshToken(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		panic(err)
	}
	var user entity.User
	filter := bson.D{{Key: "username", Value: data["username"]}}
	err := entity.UserCollection.FindOne(c.Context(), filter).Decode(&user)
	if err != nil {
		c.Status(fiber.StatusForbidden)
		return c.SendString("User not found")
	}
	return c.SendString("Hello, World!")
}

func VerifyToken(c *fiber.Ctx) error {
	cookie := c.Cookies("token")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.Status(fiber.StatusForbidden)
		return c.SendString("Unauthenticated")
	}
	return c.Next()
}

func VerifyAdmin(c *fiber.Ctx) error {
	cookie := c.Cookies("token")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.Status(fiber.StatusForbidden)
		return c.SendString("Unauthenticated")
	}
	// Retrieve user from database
	var user entity.User
	objectId, err := primitive.ObjectIDFromHex(claims["userId"].(string))
	if err != nil {
		fmt.Println("Invalid id")
	}
	filter := bson.D{{Key: "_id", Value: objectId}}
	err = entity.UserCollection.FindOne(c.Context(), filter).Decode(&user)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.SendString("User not found")
	}
	// Check if user is admin
	if user.Role != "admin" {
		c.Status(fiber.StatusForbidden)
		return c.SendString("Unauthorized Not Admin")
	}
	return c.Next()
}

func GetUser(c *fiber.Ctx) error {
	cursor, err := entity.UserCollection.Find(context.TODO(), bson.D{}, options.Find())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	var users []entity.User
	if err = cursor.All(context.Background(), &users); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(users)
}

func Logout(c *fiber.Ctx) error {
	c.ClearCookie("token")
	return c.SendStatus(fiber.StatusOK)
}

func AddUser(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var result, err = entity.UserCollection.InsertOne(c.Context(), entity.User{
		ID:       primitive.NewObjectID(),
		Username: data["username"],
		Password: data["password"],
		Role:     data["role"],
	})
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	c.SendStatus(fiber.StatusCreated)
	return c.JSON(result)
}

func DeleteUser(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	objectId, err := primitive.ObjectIDFromHex(data["id"])
	if err != nil {
		fmt.Println("Invalid id")
	}
	filter := bson.D{{Key: "_id", Value: objectId}}
	_, err = entity.UserCollection.DeleteOne(c.Context(), filter)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}
