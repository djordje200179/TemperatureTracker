package handlers

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/storage/cache"
	"net/http"
)

type Index struct {
	Context
	Readings []reading.Reading
}

func (handler Index) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handler.Readings = cache.Instance().GetLatestReadings()
	handler.UseTemplate("index", writer, handler)
}
