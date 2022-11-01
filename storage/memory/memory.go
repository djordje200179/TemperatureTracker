package memory

import "TemperatureTracker/sensor"

type memory struct {
	readings []sensor.Reading
}

func (storage *memory) AddReading(reading sensor.Reading) error {
	storage.readings = append(storage.readings, reading)

	return nil
}

func (storage *memory) GetAllReadings() ([]sensor.Reading, error) {
	return storage.readings, nil
}
