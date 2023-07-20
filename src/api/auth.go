package auth

import "github.com/gofiber/fiber/v2"

func RouterFigo() *fiber.App {
	router := fiber.New()
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Estás en Auth")
	})
	router.Get("/registrar", func(c *fiber.Ctx) error {
		return c.SendString("Registrado")
	})
	router.Get("/salir", func(c *fiber.Ctx) error {
		return c.SendString("Saliendo de la app")
	})
	return router
}
