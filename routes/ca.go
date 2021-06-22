package routes

import (
	"github.com/LfJohnVo/fiber-crud/controllers"
	"github.com/gofiber/fiber/v2"
)

func CaRoute(route fiber.Router) {
	//route.Get("", controllers.GetCa)
	route.Post("", controllers.CreateCa)
}