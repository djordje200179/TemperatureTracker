package memory

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/sensors/global"
	"TemperatureTracker/data/storage/cache"
	"sync"
)

type memory struct {
	readings []reading.Reading

	lock sync.Mutex
}

func (storage *memory) AddReading(reading reading.Reading) error {
	storage.lock.Lock()
	storage.readings = append(storage.readings, reading)
	storage.lock.Unlock()

	cache.Instance().AddReading(reading)

	return nil
}

func (storage *memory) GetReading(sensor global.Sensor) (reading.Reading, error) {
	//return storage.latestReadings[sensor], nil
	panic("implement me")
}

func (storage *memory) GetAllReadings() ([]reading.Reading, error) {
	copiedReadings := make([]reading.Reading, len(storage.readings))
	copy(copiedReadings, storage.readings)

	return copiedReadings, nil
}
