package pages

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/storage/cache"
	"TemperatureTracker/server/templates"
	"net/http"
)

type indexData struct {
	Readings []reading.Reading
}

func (context Context) Index(writer http.ResponseWriter, _ *http.Request) {
	readings := cache.Instance().GetLatestReadings()

	templates.Use("index", writer, indexData{readings})
}
