package memory

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/sensors/global"
	"sync"
)

type memory struct {
	readings       []reading.Reading
	latestReadings map[global.Sensor]reading.Reading

	lock sync.Mutex
}

func (storage *memory) AddReading(reading reading.Reading) error {
	storage.lock.Lock()
	defer storage.lock.Unlock()

	storage.readings = append(storage.readings, reading)
	storage.latestReadings[reading.Sensor] = reading

	return nil
}

func (storage *memory) GetReading(sensor global.Sensor) (reading.Reading, error) {
	return storage.latestReadings[sensor], nil
}

func (storage *memory) GetLatestReadings() ([]reading.Reading, error) {
	copiedReadings := make([]reading.Reading, 0, len(storage.latestReadings))
	for _, currentReading := range storage.latestReadings {
		copiedReadings = append(copiedReadings, currentReading)
	}

	return copiedReadings, nil
}

func (storage *memory) GetAllReadings() ([]reading.Reading, error) {
	copiedReadings := make([]reading.Reading, len(storage.readings))
	copy(copiedReadings, storage.readings)

	return copiedReadings, nil
}
