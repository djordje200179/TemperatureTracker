package cache

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/sensors/global"
	"TemperatureTracker/data/storage"
)

var instance *cache

func Instance() storage.Cache {
	if instance == nil {
		instance = &cache{
			readings: make(map[global.Sensor]reading.Reading),
		}
	}

	return instance
}
