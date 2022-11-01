package database

import (
	"TemperatureTracker/sensor"
	"TemperatureTracker/storage"
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

func (storage *database) AddReading(reading *sensor.Reading) error {
	//TODO implement me
	panic("implement me")
}

func (storage *database) GetAllReadings() ([]*sensor.Reading, error) {
	//TODO implement me
	panic("implement me")
}
