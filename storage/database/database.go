package database

import (
	"TemperatureTracker/sensor"
	"database/sql"
)

type database sql.Conn

func (storage *database) AddReading(reading *sensor.Reading) error {
	//TODO implement me
	panic("implement me")
}

func (storage *database) GetAllReadings() ([]*sensor.Reading, error) {
	//TODO implement me
	panic("implement me")
}
