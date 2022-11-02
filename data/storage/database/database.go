package database

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/sensors/global"
	"TemperatureTracker/data/storage"
	"TemperatureTracker/data/storage/cache"
	"database/sql"
	_ "modernc.org/sqlite"
)

type database sql.DB

func Open(path string) (storage.Storage, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return (*database)(db), nil
}

func (storage *database) AddReading(reading reading.Reading) error {
	//TODO implement me
	panic("implement me")

	cache.Instance().AddReading(reading)

	return nil
}

func (storage *database) GetReading(sensor global.Sensor) (reading.Reading, error) {
	//TODO implement me
	panic("implement me")
}

func (storage *database) GetAllReadings() ([]reading.Reading, error) {
	//TODO implement me
	panic("implement me")
}
