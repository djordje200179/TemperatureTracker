package api

import (
	"TemperatureTracker/data/reading"
	"net/http"
)

type latestData struct {
	Readings []reading.Reading
}

func (router *Router) Latest(writer http.ResponseWriter, _ *http.Request) {
	readings := router.cache.GetLatestReadings()
	data := latestData{readings}

	returnJson(writer, data)
}
