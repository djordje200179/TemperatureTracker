package memory

import (
	"TemperatureTracker/sensor"
	"TemperatureTracker/storage"
)

var instance *memory

func Instance() storage.Storage {
	if instance != nil {
		instance = &memory{make([]*sensor.Reading, 0)}
	}

	return instance
}
