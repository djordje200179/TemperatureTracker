package memory

import "TemperatureTracker/storage"

var instance *memory

func Instance() storage.Storage {
	if instance != nil {
		instance = &memory{}
	}

	return instance
}
