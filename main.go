package main

import (
	"fmt"
	"github.com/LfJohnVo/fiber-crud/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	fmt.Printf("TEST")
	app := fiber.New()

	//Middleware que ademas de restringir, permite observar los movimientos dentro de la consola
	app.Use(logger.New()) // new
	setupRoutes(app)
	err := app.Listen(":8000")

	if err != nil {
		panic(err)
	}

}

func setupRoutes(app *fiber.App) {
	//Sirve el endpoint "/"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":  true,
			"message": "Bienvenido 😉",
		})
	})

	//Grupo api
	api := app.Group("/api")

	//Sirve los endpoint a partir de /api
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Esta api se muestra desde /api 😉",
		})
	})

	// connect todo routes
	routes.TodoRoute(api.Group("/todos"))

}
