package main

import (
	"TemperatureTracker/data/sensors/local"
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
	local.RegisterDS18B20()

	local.StartLogging(storage, loggingInterval)

	err := server.Start(storage, port)
	if err != nil {
		log.Fatal(err)
	}
}
