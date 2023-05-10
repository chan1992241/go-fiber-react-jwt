package main

import (
	"chan1992241/backend/cmd/config"
	"chan1992241/backend/cmd/controller"
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func TestLogin(t *testing.T) {
	app := fiber.New()
	app.Use(cors.New(cors.Config(config.CorsConfig)))
	app.Post("/login", controller.Login)

	req, error := http.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"username":"jinyee","password":"123"}`))
	if error != nil {
		t.Error(error)
	}
	req.Header.Set("Content-Type", "application/json")

	response, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	// t.Log(response)

	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code is not 200. Got %d", response.StatusCode)
	}
}

// func TestGetUsers(t *testing.T) {
// 	app := fiber.New()
// 	app.Use(cors.New(cors.Config(config.CorsConfig)))
// 	//Login first to get token and cookie
// 	app.Post("/login", controller.Login)

// 	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"username":"jinyee","password":"123"}`))
// 	req.Header.Set("Content-Type", "application/json")

// 	response, err := app.Test(req)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Log(response.Cookies())

// 	// app.Get("/users", controller.VerifyToken, controller.VerifyAdmin, controller.GetUser)

// 	// req, _ = http.NewRequest(http.MethodGet, "/users", nil)
// 	// req.Header.Set("Content-Type", "application/json")

// 	// response, err = app.Test(req)
// 	// t.Log(response)

// 	// if err != nil {
// 	// 	t.Error(err)
// 	// }
// 	// if response.StatusCode != http.StatusOK {
// 	// 	t.Errorf("Status code is not 200. Got %d", response.StatusCode)
// 	// }
// }
