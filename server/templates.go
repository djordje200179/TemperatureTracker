package server

import (
	templateDir "TemperatureTracker/templates"
	"html/template"
	"io/fs"
	"path/filepath"
	"strings"
)

type TemplateMap = map[string]*template.Template

func loadTemplates() (TemplateMap, error) {
	files, err := fs.ReadDir(templateDir.FS, ".")
	if err != nil {
		return nil, err
	}

	templates := make(TemplateMap)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		templateName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		templates[templateName] = template.Must(template.ParseFS(templateDir.FS, fileName))
	}

	return templates, nil
}
