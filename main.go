package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	fmt.Printf("TEST")
	app := fiber.New()

	//Middleware que ademas de restringir, permite observar los movimientos dentro de la consola
	app.Use(logger.New()) // new
	//Sirve el endpoint "/"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success":  true,
			"message": "Bienvenido 😉",
		})
	})

	err := app.Listen(":8000")

	if err != nil {
		panic(err)
	}
}
