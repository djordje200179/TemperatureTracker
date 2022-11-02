package memory

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/sensors/global"
	"TemperatureTracker/data/storage"
)

var instance *memory

func Instance() storage.Storage {
	if instance == nil {
		instance = &memory{
			readings:       make([]reading.Reading, 0),
			latestReadings: make(map[global.Sensor]reading.Reading),
		}
	}

	return instance
}
