package config

import "github.com/gofiber/fiber/v2"

var allowOrigin []string = []string{"http://localhost:5173"}

type Config struct {
	Next             func(c *fiber.Ctx) bool
	AllowOrigins     string
	AllowMethods     string
	AllowHeaders     string
	AllowCredentials bool
	ExposeHeaders    string
	MaxAge           int
}

var CorsConfig = Config{
	AllowOrigins:     allowOrigin[0],
	AllowCredentials: true,
	AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
}
