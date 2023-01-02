package main

import (
	// "github.com/gofiber/fiber"
	"github.com/infuman69/expense-tracker-api/database"
)

// func setUpRoutes(app *fiber.App) {

// }
func main() {
	// app := fiber.New()
	// setUpRoutes(app)
	database.Database()
}
