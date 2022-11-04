package logger

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/storage"
	"time"
)

type Logger struct {
	Sensor  reading.Sensor
	Storage storage.Storage

	*time.Ticker
}

func Start(sensor reading.Sensor, storage storage.Storage, period time.Duration) Logger {
	logger := Logger{
		Sensor:  sensor,
		Storage: storage,
		Ticker:  time.NewTicker(period),
	}

	go logger.Log()

	return logger
}

func (logger Logger) Log() {
	for range logger.C {
		err := logger.SingleLog()
		if err != nil {
			panic(err)
		}
	}
}

func (logger Logger) SingleLog() error {
	data, err := logger.Sensor.Read()
	if err != nil {
		return err
	}

	err = logger.Storage.AddReading(data)
	if err != nil {
		return err
	}

	return nil
}
