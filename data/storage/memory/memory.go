package memory

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/storage/cache"
	"sync"
)

type memory struct {
	readings []reading.Reading

	lock sync.RWMutex
}

func (storage *memory) AddReading(reading reading.Reading) error {
	storage.lock.Lock()
	storage.readings = append(storage.readings, reading)
	storage.lock.Unlock()

	cache.Instance().AddReading(reading)

	return nil
}

func (storage *memory) GetReadings(sensor reading.Sensor) ([]reading.Reading, error) {
	storage.lock.RLock()
	defer storage.lock.RUnlock()

	filteredReadings := make([]reading.Reading, 0)
	for _, currentReading := range storage.readings {
		if currentReading.Sensor == sensor {
			filteredReadings = append(filteredReadings, currentReading)
		}
	}

	return filteredReadings, nil
}

func (storage *memory) GetAllReadings() ([]reading.Reading, error) {
	storage.lock.RLock()
	defer storage.lock.RUnlock()

	copiedReadings := make([]reading.Reading, len(storage.readings))
	copy(copiedReadings, storage.readings)

	return copiedReadings, nil
}
