package routes

import (
	"github.com/LfJohnVo/fiber-crud/controllers"
	"github.com/gofiber/fiber/v2"
)

//Ruta uno, get all
func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
}
