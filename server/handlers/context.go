package handlers

import (
	"TemperatureTracker/data/storage"
	"net/http"
)

type Context struct {
	Storage storage.Storage
}

func (context Context) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", context.Index)
}
