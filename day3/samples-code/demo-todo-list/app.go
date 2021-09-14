package main

import (
	"demo-todo-list/config"
	"demo-todo-list/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/joho/godotenv" // new
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint",
		})
	})

	api := app.Group("/api")
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api point",
		})
	})

	routes.TodoRoute(api.Group("/todos"))
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	os.Setenv("MONGO_URI", "mongodb+srv://golang-demo:golang-demo@golang-demo.vd9n5.mongodb.net/golang-demo?retryWrites=true&w=majority")
	// config db
	config.ConnectDB()

	// setup routes
	setupRoutes(app)

	// Listen on server 8000 and catch error if any
	err = app.Listen(":8000")

	// handle error
	if err != nil {
		panic(err)
	}

}
