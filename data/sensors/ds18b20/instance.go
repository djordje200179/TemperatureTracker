package ds18b20

import (
	"TemperatureTracker/data/reading"
)

var instance ds18b20 = ds18b20{
	Id: findId(),
}

func Instance() reading.Sensor {
	return instance
}
