package handlers

import (
	"TemperatureTracker/data/storage"
	"TemperatureTracker/server/templates"
	"net/http"
)

type Context struct {
	Storage storage.Storage

	Templates templates.Map
}

func MakeContext(storage storage.Storage) Context {
	return Context{
		Storage:   storage,
		Templates: templates.Load(),
	}
}

func (context Context) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", context.Index)
}
