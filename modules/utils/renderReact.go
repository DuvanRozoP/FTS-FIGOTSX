package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"rogchap.com/v8go"
)

func RenderReact(ctx *v8go.Context, url string) string {

	reactAppArgs := map[string]interface{}{
		"url": url,
	}
	appArgs, err := json.Marshal(reactAppArgs)
	if err != nil {
		fmt.Println("Error al Json.Marshal => ", err)
	}
	runRenderReactApp := fmt.Sprintf(`var render = FIGOTSX.render(%s)`, appArgs)
	_, err = ctx.RunScript(runRenderReactApp, "render.js")
	if err != nil {
		fmt.Println("Error de codigo JS:", err)
		os.Exit(1)
	}

	getHtml, err := ctx.RunScript("render.html", "values.js")
	if err != nil {
		fmt.Println("Error de codigo JS:", err)
		os.Exit(1)
	}

	return getHtml.String()
}
