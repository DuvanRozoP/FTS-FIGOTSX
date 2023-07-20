package router

import (
	auth "github.com/DuvanRozoP/FTS-FIGOTSX/src/api"
	"github.com/gofiber/fiber/v2"
)

func RouterApp(app *fiber.App) {
	app.Mount("/auth", auth.RouterFigo())

}
