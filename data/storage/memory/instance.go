package memory

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/storage"
)

var instance = memory{
	readings: make([]reading.Reading, 0),
}

func Instance() storage.Storage {
	return &instance
}
