package handlers

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/storage/cache"
	"net/http"
)

type indexData struct {
	Readings []reading.Reading
}

func (context Context) Index(writer http.ResponseWriter, _ *http.Request) {
	readings := cache.Instance().GetLatestReadings()

	context.UseTemplate("index", writer, indexData{readings})
}
