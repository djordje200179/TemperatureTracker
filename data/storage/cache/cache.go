package cache

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/sensors/global"
	"sync"
)

type cache struct {
	readings map[global.Sensor]reading.Reading
	lock     sync.RWMutex
}

func (cache *cache) AddReading(newReading reading.Reading) {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	cache.readings[newReading.Sensor] = newReading
}

func (cache *cache) GetLatestReadings() []reading.Reading {
	cache.lock.RLock()
	defer cache.lock.RUnlock()

	clonedReadings := make([]reading.Reading, len(cache.readings))
	for _, currentReading := range cache.readings {
		clonedReadings = append(clonedReadings, currentReading)
	}

	return clonedReadings
}
