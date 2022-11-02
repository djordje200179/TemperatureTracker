package templates

import (
	"embed"
	"html/template"
	"io/fs"
	"path/filepath"
	"strings"
)

//go:embed *.html
var filesystem embed.FS

type Map = map[string]*template.Template

func Load() (Map, error) {
	files, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}

	templates := make(Map)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		templateName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		templates[templateName] = template.Must(template.ParseFS(filesystem, fileName))
	}

	return templates, nil
}
