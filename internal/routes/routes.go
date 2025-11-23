package routes

import (
	"GoAndReact/internal/handlers"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/todos", handlers.GetTodos)
	api.Post("/todos", handlers.CreateTodo)
	api.Patch("/todos/:id", handlers.UpdateTodo)
	api.Delete("/todos/:id", handlers.DeleteTodo)
}