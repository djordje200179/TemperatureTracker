package handlers

import (
	"TemperatureTracker/data/reading"
	"fmt"
	"net/http"
)

type IndexData struct {
	Readings []reading.Reading
}

func (context Context) Index(writer http.ResponseWriter, request *http.Request) {
	var err error

	latestReadings, err := context.Storage.GetLatestReadings()
	if err != nil {
		fmt.Println(err)
	}

	data := IndexData{Readings: latestReadings}

	err = context.Templates["index"].Execute(writer, data)
	if err != nil {
		fmt.Println(err)
	}
}
