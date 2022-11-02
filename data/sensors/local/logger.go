package local

import (
	"TemperatureTracker/data/storage"
	"fmt"
	"time"
)

type Logger struct {
	Storage storage.Storage
}

func StartLogging(storage storage.Storage, period time.Duration) {
	logger := Logger{Storage: storage}
	go logger.LogTemperatures(period)
}

func (logger Logger) LogTemperatures(period time.Duration) {
	for range time.Tick(period) {
		logger.LogTemperature()
	}
}

func (logger Logger) LogTemperature() {
	for _, currSensor := range Sensors() {
		reading, err := currSensor.Read()
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = logger.Storage.AddReading(reading)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
