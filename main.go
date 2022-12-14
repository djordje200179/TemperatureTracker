package main

import (
	"TemperatureTracker/cli"
	"TemperatureTracker/data/logger"
	"TemperatureTracker/data/sensors/ds18b20"
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

	loggingInterval = 5 * time.Minute
)

func main() {
	logger.Start(ds18b20.Instance(), storage, loggingInterval)

	cli.Default().Storage = storage
	go cli.Default().Handle()

	err := server.New(storage).Start(port)
	if err != nil {
		log.Fatal(err)
	}
}
