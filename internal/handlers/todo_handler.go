package handlers

import (
	"context"
	"GoAndReact/internal/database"
	"GoAndReact/internal/models"

	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodos(c fiber.Ctx) error {
	var todos []models.Todo

	cursor, err := database.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch todos"})
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to decode todo"})
		}
		todos = append(todos, todo)
	}

	return c.JSON(todos)
}

func CreateTodo(c fiber.Ctx) error {
	todo := new(models.Todo)

	if err := c.Bind().Body(todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Todo body cannot be empty"})
	}

	todo.Completed = false

	result, err := database.Collection.InsertOne(context.Background(), todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create todo"})
	}

	todo.ID = result.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(todo)
}

func UpdateTodo(c fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	update := bson.M{"$set": bson.M{"completed": true}}
	_, err = database.Collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update todo"})
	}

	return c.JSON(fiber.Map{"success": true})
}

func DeleteTodo(c fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	_, err = database.Collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete todo"})
	}

	return c.JSON(fiber.Map{"success": true})
}