package main

import (
	"chan1992241/backend/cmd/config"
	"chan1992241/backend/cmd/controller"
	mongodb "chan1992241/backend/cmd/model/bean"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func load_env() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	load_env()
	mongodb.MongodbInitialization()
	app := fiber.New()
	app.Use(cors.New(cors.Config(config.CorsConfig)))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Test!")
	})
	app.Post("/register", controller.Register)
	app.Post("/login", controller.Login)
	app.Get("/users", controller.VerifyToken, controller.VerifyAdmin, controller.GetUser)
	app.Get("/logout", controller.VerifyToken, controller.Logout)
	app.Get("/refresh", controller.RefreshToken)
	app.Post("/add", controller.VerifyToken, controller.VerifyAdmin, controller.AddUser)
	app.Delete("/delete", controller.VerifyToken, controller.VerifyAdmin, controller.DeleteUser)
	app.Listen(":3000")
}
