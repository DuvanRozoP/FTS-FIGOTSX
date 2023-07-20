package main

import (
	router "github.com/DuvanRozoP/FTS-FIGOTSX/modules/lib"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Static("/", "./public")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	router.RouterApp(app)
	app.Listen(":3000")
}
