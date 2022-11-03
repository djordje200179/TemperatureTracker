package handlers

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/storage"
	"TemperatureTracker/data/storage/cache"
	"TemperatureTracker/server/templates"
	"net/http"
)

type Context struct {
	Storage storage.Storage
}

func (context Context) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", context.Index)
}

type indexData struct {
	Readings []reading.Reading
}

func (context Context) Index(writer http.ResponseWriter, _ *http.Request) {
	readings := cache.Instance().GetLatestReadings()

	templates.Use("index", writer, indexData{readings})
}
