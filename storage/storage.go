package storage

import "TemperatureTracker/sensor"

type Storage interface {
	AddReading(reading *sensor.Reading) error

	GetAllReadings() ([]*sensor.Reading, error)
}
