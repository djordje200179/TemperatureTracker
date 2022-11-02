package logger

import (
	"TemperatureTracker/sensor"
	"TemperatureTracker/storage"
	"fmt"
	"time"
)

type Logger struct {
	Sensors []sensor.Sensor
	Storage storage.Storage
}

func Start(storage storage.Storage, period time.Duration) error {
	sensors, err := sensor.Sensors()
	if err != nil {
		return err
	}

	logger := Logger{Sensors: sensors, Storage: storage}
	go logger.LogTemperatures(period)

	return nil
}

func (logger Logger) LogTemperatures(period time.Duration) {
	for range time.Tick(period) {
		logger.LogTemperature()
	}
}

func (logger Logger) LogTemperature() {
	for _, currSensor := range logger.Sensors {
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
