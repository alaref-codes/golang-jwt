package main

import (
	"log"

	"github.com/alaref-codes/basic-auth/database"
	"github.com/alaref-codes/basic-auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	routes.SetUpRoutes(app)
	database.InitDatabase()
	app.Use(logger.New())
	log.Fatal(app.Listen(":3000"))
}
