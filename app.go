package main

import (
	"log"

	"github.com/alaref-codes/basic-auth/database"
	"github.com/alaref-codes/basic-auth/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"

	jwtware "github.com/gofiber/jwt/v3"

	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	setUpRoutes(app)
	initDatabase()
	log.Fatal(app.Listen(":3000"))
}

func initDatabase() {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	database.DBConn = db
	if err != nil {
		panic("failed to connect database")
	}
	database.DBConn.AutoMigrate(&user.User{})
}

func setUpRoutes(app *fiber.App) {
	app.Post("/users/signin", user.Signin)
	app.Post("/users", user.CreateUser) // Sign up page

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	// Restricted routes !
	app.Get("/users", user.GetUsers)

	app.Delete("/users/:email", user.DeleteUser)

}
