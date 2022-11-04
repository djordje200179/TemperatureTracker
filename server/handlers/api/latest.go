package api

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/storage/cache"
	"encoding/json"
	"fmt"
	"net/http"
)

type latestData struct {
	Readings []reading.Reading
}

func (context Context) Latest(writer http.ResponseWriter, _ *http.Request) {
	fmt.Println("Latest")

	readings := cache.Instance().GetLatestReadings()

	jsonData, err := json.Marshal(latestData{readings})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(jsonData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
