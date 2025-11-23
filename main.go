package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"GoAndReact/config"
	"GoAndReact/internal/database"
	"GoAndReact/internal/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	app := fiber.New()

	client := database.Connect()
	database.SetCollection(client)

	config.SetupCORS(app)

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on :%s", port)
	log.Fatal(app.Listen("0.0.0.0:" + port))
}