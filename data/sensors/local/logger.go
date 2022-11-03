package local

import (
	"TemperatureTracker/data/storage"
	"fmt"
	"time"
)

type Logger struct {
	Storage storage.Storage
	Period  time.Duration
}

func StartLogging(storage storage.Storage, period time.Duration) {
	logger := Logger{
		Storage: storage,
		Period:  period,
	}

	go logger.ContinuousLogging()
}

func (logger Logger) ContinuousLogging() {
	for range time.Tick(logger.Period) {
		for _, currSensor := range Sensors() {
			go logger.SingleSensorLog(currSensor)
		}
	}
}

func (logger Logger) SingleSensorLog(sensor Sensor) {
	reading, err := sensor.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = logger.Storage.AddReading(reading)
	if err != nil {
		fmt.Println(err)
		return
	}
}
