package database

import "TemperatureTracker/storage"

var instance *database

func Instance() storage.Storage {
	if instance == nil {
		instance = new(database)
	}

	return instance
}
