package api

import "github.com/gofiber/fiber/v2"

func Index() *fiber.App {
	router := fiber.New()
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Estás en Auth v3")
	})

	router.Get("/registrar", func(c *fiber.Ctx) error {
		return c.SendString("Registrado v3")
	})

	router.Get("/salir", func(c *fiber.Ctx) error {
		return c.SendString("Saliendo de la app v3")
	})
	return router
}
