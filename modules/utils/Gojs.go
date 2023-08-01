package utils

import (
	"fmt"
	"os"

	"rogchap.com/v8go"
)

func createGlobalObject(iso *v8go.Isolate) *v8go.ObjectTemplate {
	goApp := v8go.NewObjectTemplate(iso)

	global := v8go.NewObjectTemplate(iso)
	global.Set("FIGOTSX", goApp)
	return global
}

func GoJs(appBundlePath string) (*v8go.Context, error) {
	iso := v8go.NewIsolate()
	global := createGlobalObject(iso)
	ctx := v8go.NewContext(iso, global)

	appBundle, err := ReadFile(appBundlePath)
	if err != nil {
		fmt.Println("Error al leer el archivo bundle:", err)
		os.Exit(1)
	}

	_, err = ctx.RunScript(string(appBundle), "client.js")
	if err != nil {
		fmt.Println("Error al ejecutar el archivo bundle:", err)
		os.Exit(1)
	}

	return ctx, nil
}
