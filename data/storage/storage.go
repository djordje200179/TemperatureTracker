package storage

import (
	"TemperatureTracker/data/reading"
)

type Storage interface {
	AddReading(reading reading.Reading) error

	GetReadings(sensor reading.Sensor) ([]reading.Reading, error)
	GetAllReadings() ([]reading.Reading, error)
}

type Cache interface {
	AddReading(reading reading.Reading)

	GetLatestReadings() []reading.Reading
}
