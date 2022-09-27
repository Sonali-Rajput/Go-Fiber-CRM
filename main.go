package main

import (
	"fmt"

	"github.com/Sonali-Rajput/go-fiber-crm/database"
	"github.com/Sonali-Rajput/go-fiber-crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBconn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connection opened to database")
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")

}

func main() {
	app := fiber.New() //instance of fiber
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBconn.Close()
}
