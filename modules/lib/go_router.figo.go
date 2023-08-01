package router

import (
	cape1 "github.com/DuvanRozoP/FTS-FIGOTSX/src/api"
	
	"github.com/gofiber/fiber/v2"
)

func RouterApp(app *fiber.App) {
	api := app.Group("/api")
	api.Mount("/", cape1.Index())
	
}