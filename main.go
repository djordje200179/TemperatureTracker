package main

import (
	"TemperatureTracker/sensor"
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	sensors, err := sensor.Sensors()
	if err != nil {
		panic(err)
	}

	mainSensor := sensors[0]
	temp, err := mainSensor.Read()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Cao Jano!\nKada si ti tu, zaista je vruce...\nEvo sad je paklenih %s...", temp)
}

func main() {
	http.HandleFunc("/", hello)

	http.ListenAndServe(":8090", nil)
}
