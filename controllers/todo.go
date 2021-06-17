package controllers

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
)

//Modelo todo
type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

//Modelo de ejemplo
var todos = []*Todo{
	{
		Id:        1,
		Title:     "Walk the dog 🦮",
		Completed: false,
	},
	{
		Id:        2,
		Title:     "Walk the cat 🐈",
		Completed: false,
	},
}

//Get all todos
func GetTodos(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todos": todos,
		},
	})
}