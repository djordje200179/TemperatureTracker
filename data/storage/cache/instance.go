package cache

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/storage"
)

var instance = cache{
	readings: make(map[reading.Sensor]reading.Reading),
}

func Instance() storage.Cache {
	return &instance
}
