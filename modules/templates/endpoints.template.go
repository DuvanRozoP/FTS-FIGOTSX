package templates

import (
	"strings"
	"text/template"
)

func TEndpoints(importModules string, moutRouter string) string {
	templateText := `package router

import (
	{{.ImportModules}}
	"github.com/gofiber/fiber/v2"
)

func RouterApp(app *fiber.App) {
	{{.MountRouter}}
}`

	data := struct {
		ImportModules string
		MountRouter   string
	}{
		ImportModules: importModules,
		MountRouter:   moutRouter,
	}

	tmpl := template.Must(template.New("endpoints").Parse(templateText))

	var builder strings.Builder
	err := tmpl.Execute(&builder, data)
	if err != nil {
		// Manejar el error si ocurrió
		return ""
	}

	return builder.String()
}
