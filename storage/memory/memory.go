package memory

import (
	"TemperatureTracker/sensor"
	"sync"
)

type memory struct {
	readings []*sensor.Reading
	lock     sync.Mutex
}

func (storage *memory) AddReading(reading *sensor.Reading) error {
	storage.lock.Lock()
	defer storage.lock.Unlock()

	storage.readings = append(storage.readings, reading)

	return nil
}

func (storage *memory) GetAllReadings() ([]*sensor.Reading, error) {
	storage.lock.Lock()
	defer storage.lock.Unlock()

	return storage.readings, nil
}
