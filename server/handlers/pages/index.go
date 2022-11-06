package pages

import (
	"TemperatureTracker/data/reading"
	"net/http"
)

type indexData struct {
	Readings []reading.Reading
}

func (router *Router) Index(writer http.ResponseWriter, _ *http.Request) {
	readings := router.Cache.GetLatestReadings()
	data := indexData{readings}

	returnPage("index", writer, data)
}
