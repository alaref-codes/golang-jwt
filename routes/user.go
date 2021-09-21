package routes

import (
	"github.com/alaref-codes/basic-auth/services"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func SetUpRoutes(app *fiber.App) {
	app.Post("/users/signin", services.Signin)
	app.Post("/users", services.CreateUser) // Sign up page

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	// Restricted routes !
	app.Get("/users", services.GetUsers)

	app.Delete("/users/:email", services.DeleteUser)

}
