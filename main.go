package main

import (
	"TemperatureTracker/logger"
	"TemperatureTracker/server"
	"TemperatureTracker/storage/memory"
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
	var err error

	err = logger.Start(storage, loggingInterval)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Start(storage, port)
	if err != nil {
		log.Fatal(err)
	}
}
