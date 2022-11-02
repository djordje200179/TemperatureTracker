package storage

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/sensors/global"
)

type Storage interface {
	AddReading(reading reading.Reading) error
	GetReading(sensor global.Sensor) (reading.Reading, error)

	GetAllReadings() ([]reading.Reading, error)
}

type Cache interface {
	AddReading(reading reading.Reading)

	GetLatestReadings() []reading.Reading
}
