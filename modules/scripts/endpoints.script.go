package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"sync"

	"github.com/DuvanRozoP/FTS-FIGOTSX/modules/templates"
)

func main() {
	directory := "src/api/"
	absPath, err := filepath.Abs(directory)
	if err != nil {
		log.Fatalf("Error obtaining absolute path: %v", err)
	}

	files, err := ioutil.ReadDir(absPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	var goFiles []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".go" {
			fileName := strings.TrimSuffix(file.Name(), ".go")
			goFiles = append(goFiles, fileName)
		}
	}

	var importModules strings.Builder
	var mountRouters strings.Builder
	var wg sync.WaitGroup
	wg.Add(len(goFiles))

	for _, file := range goFiles {
		go func(fileName string) {
			defer wg.Done()
			importModule := fileName + " \"github.com/DuvanRozoP/go-fiber/src/api\""
			mountRouter := "app.Mount(\"/" + fileName + "\", " + fileName + ".RouterFigo())"

			importModules.WriteString(importModule)
			importModules.WriteString("\n")
			mountRouters.WriteString(mountRouter)
			mountRouters.WriteString("\n")
		}(file)
	}

	wg.Wait()

	content := templates.TEndpoints(importModules.String(), mountRouters.String())

	err = ioutil.WriteFile("modules/lib/router.figo.go", []byte(content), 0644)
	if err != nil {
		log.Fatalf("Error writing file: %v", err)
	}

	log.Println("Create endpoints")
}
