package routes

import (
	"github.com/LfJohnVo/fiber-crud/controllers"
	"github.com/gofiber/fiber/v2"
)

func CsrRoute(route fiber.Router) {
	route.Get("", controllers.GetCsr)
	route.Post("", controllers.CreateCsr)
}
