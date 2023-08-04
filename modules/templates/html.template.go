package templates

import (
	"text/template"
)

func THtml() *template.Template {
	tmpl := template.Must(template.New("index").Parse(`<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" href="/index.bundle.css">
	<title>{{.Title}}</title>
</head>
<body>
	<div id="root">{{.Html}}</div>
</body>
</html>`))

	return tmpl
}
