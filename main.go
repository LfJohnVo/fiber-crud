package main

import (
	"fmt"
	_ "github.com/LfJohnVo/fiber-crud/docs"
	"github.com/LfJohnVo/fiber-crud/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

func main() {
	fmt.Printf("TEST")
	app := fiber.New()

	//Middleware que ademas de restringir, permite observar los movimientos dentro de la consola
	app.Use(logger.New()) // new

	app.Use(recover.New()) // recover from crash

	setupRoutes(app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})


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

	// connect csr routes
	routes.CsrRoute(api.Group("/csr"))
	// connect ca routes
	routes.CaRoute(api.Group("/ca"))

}
