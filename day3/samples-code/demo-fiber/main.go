package main

import (
	"demo-fiber/book"
	"demo-fiber/database"
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

//https://tutorialedge.net/golang/basic-rest-api-go-fiber/
func main() {
	//demoFiber1()
	demoFiber2()
}

func demoFiber1() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Listen(8080)
}

func demoFiber2() {
	app := fiber.New()
	initDatabase()
	setRoutes(app)
	app.Listen(8080)
	defer database.DBConn.Close()
}

func setRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}
