package server

import (
	"TemperatureTracker/storage"
	"fmt"
	"net/http"
)

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
