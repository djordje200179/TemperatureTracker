package server

import (
	"TemperatureTracker/storage"
	"TemperatureTracker/templates"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
)

type TemplateMap = map[string]*template.Template

func loadTemplates() (TemplateMap, error) {
	files, err := fs.ReadDir(templates.FS, ".")
	if err != nil {
		return nil, err
	}

	templateMap := make(TemplateMap)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		templateName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		templateMap[templateName] = template.Must(template.ParseFS(templates.FS, fileName))
	}

	return templateMap, nil
}

type Context struct {
	Storage   storage.Storage
	Templates TemplateMap
}

func MakeContext(storage storage.Storage) (Context, error) {
	templates, err := loadTemplates()
	if err != nil {
		return Context{}, err
	}

	return Context{Storage: storage, Templates: templates}, nil
}

func (context Context) Index(w http.ResponseWriter, req *http.Request) {
	err := context.Templates["index"].Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
