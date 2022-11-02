package handlers

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/storage/cache"
	"fmt"
	"net/http"
)

type IndexData struct {
	Readings []reading.Reading
}

func (context Context) Index(writer http.ResponseWriter, request *http.Request) {
	data := IndexData{Readings: cache.Instance().GetLatestReadings()}

	err := context.Templates["index"].Execute(writer, data)
	if err != nil {
		fmt.Println(err)
	}
}
