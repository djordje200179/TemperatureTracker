package main

import (
	"TemperatureTracker/data/logger"
	"TemperatureTracker/data/sensors/ds18b20"
	"TemperatureTracker/data/storage/memory"
	"TemperatureTracker/server"
	"TemperatureTracker/server/cli"
	"log"
	"time"
)

var (
	storage = memory.Instance()
)

const (
	port = 8080

	loggingInterval = 5 * time.Minute
)

func main() {
	logger.Start(ds18b20.Instance(), storage, loggingInterval)

	cli.Default().Storage = storage

	err := server.Start(storage, port)
	if err != nil {
		log.Fatal(err)
	}
}
