package templates

import (
	"embed"
	"fmt"
	"html/template"
)

//go:embed *.html
var filesystem embed.FS

var templates = make(map[string]*template.Template)

func Get(templateName string) *template.Template {
	tmpl, ok := templates[templateName]
	if !ok {
		fileName := fmt.Sprintf("%s.html", templateName)

		tmpl = template.Must(template.ParseFS(filesystem, fileName))
		templates[templateName] = tmpl
	}

	return tmpl
}
