package main

import (
	"TemperatureTracker/data/sensors/local"
	"TemperatureTracker/data/sensors/local/ds18b20"
	"TemperatureTracker/data/storage/memory"
	"TemperatureTracker/server"
	"log"
	"time"
)

var (
	storage = memory.Instance()
)

const (
	port = 8080

	loggingInterval = 1 * time.Minute
)

func main() {
	local.RegisterSensor(ds18b20.Instance())

	local.StartLogging(storage, loggingInterval)

	err := server.Start(storage, port)
	if err != nil {
		log.Fatal(err)
	}
}
