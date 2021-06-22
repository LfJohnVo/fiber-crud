package routes

import (
	"github.com/LfJohnVo/fiber-crud/controllers"
	"github.com/gofiber/fiber/v2"
)

func KeysRoute(route fiber.Router) {
	//route.Get("", controllers.GetCa)
	route.Post("", controllers.CreateKeys)
}