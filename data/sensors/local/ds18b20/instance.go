package ds18b20

import "TemperatureTracker/data/sensors/local"

var instance ds18b20

func Instance() local.Sensor {
	return instance
}
