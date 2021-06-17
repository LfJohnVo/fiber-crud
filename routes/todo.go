package routes

import (
	"github.com/LfJohnVo/fiber-crud/controllers"
	"github.com/gofiber/fiber/v2"
)

//Ruta Todos
func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
	route.Post("", controllers.CreateTodo)
	route.Put("/:id", controllers.UpdateTodo)
	route.Delete("/:id", controllers.DeleteTodo)
	route.Get("/:id", controllers.GetTodo)
}


