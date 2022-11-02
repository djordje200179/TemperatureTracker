package cache

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/sensors/global"
)

type cache struct {
	readings map[global.Sensor]reading.Reading
}

func (cache *cache) AddReading(newReading reading.Reading) {
	cache.readings[newReading.Sensor] = newReading
}

func (cache *cache) GetLatestReadings() []reading.Reading {
	clonedReadings := make([]reading.Reading, len(cache.readings))
	for _, currentReading := range cache.readings {
		clonedReadings = append(clonedReadings, currentReading)
	}

	return clonedReadings
}
