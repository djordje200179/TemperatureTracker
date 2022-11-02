package handlers

import (
	"TemperatureTracker/sensor"
	"fmt"
	"net/http"
)

type IndexData struct {
	Readings []sensor.Reading
}

func (context Context) Index(writer http.ResponseWriter, request *http.Request) {
	readings := make([]sensor.Reading, 0, len(context.Sensors))
	for _, sensor := range context.Sensors {
		reading, err := sensor.Read()
		if err != nil {
			fmt.Println(err)
		}

		readings = append(readings, reading)
	}

	data := IndexData{Readings: readings}
	err := context.Templates["index"].Execute(writer, data)
	if err != nil {
		fmt.Println(err)
	}
}
