package memory

import (
	"TemperatureTracker/sensor"
	"TemperatureTracker/storage"
)

var instance *memory

func Instance() storage.Storage {
	if instance != nil {
		instance = &memory{readings: make([]*sensor.Reading, 0)}
	}

	return instance
}
