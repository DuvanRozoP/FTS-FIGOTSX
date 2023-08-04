package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	Router "github.com/DuvanRozoP/FTS-FIGOTSX/modules/lib"
	Templates "github.com/DuvanRozoP/FTS-FIGOTSX/modules/templates"
	Utils "github.com/DuvanRozoP/FTS-FIGOTSX/modules/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	Env "github.com/joho/godotenv"
)

func main() {
	currentPath, errCurrentPath := filepath.Abs(".")
	if errCurrentPath != nil {
		fmt.Println("Error al obtener la ruta absoluta:", errCurrentPath)
		os.Exit(1)
	}

	err := Env.Load()
	if err != nil {
		fmt.Println("Error al cargar el archivo .env:", err)
		os.Exit(1)
	}

	PORT := os.Getenv("PORT")
	ISDEV := os.Getenv("isDev")

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	appBundlePath := filepath.Join(currentPath, "./client/main.bundle.js")
	cssBundlePath := filepath.Join(currentPath, "./static")
	if ISDEV == "true" {
		app.Get("/metrics", monitor.New())
		app.Use(logger.New(logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		}))
		appBundlePath = filepath.Join(currentPath, "./src/cache/main.bundle.js")
		cssBundlePath = filepath.Join(currentPath, "./public")
	}

	ctx, err := Utils.GoJs(appBundlePath)
	if err != nil {
		fmt.Println("Error con GoJs v8go => ", err)
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:4001",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Static("/", cssBundlePath)
	Router.RouterApp(app)
	app.Get("/*", func(c *fiber.Ctx) error {
		currentURL := c.Path()
		contentHtml := Utils.RenderReact(ctx, currentURL)
		c.Response().Header.SetContentType(fiber.MIMETextHTML)
		return Templates.THtml().Execute(c.Response().BodyWriter(), fiber.Map{
			"Title": "FTS-FIGOTSX",
			"Html":  contentHtml,
		})
	})
	log.Println("Servidor iniciado en http://localhost:", PORT)
	err = app.Listen(":" + PORT)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %s", err)
	}
}
