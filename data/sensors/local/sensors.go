package local

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/sensors/global"
	"fmt"
)

type Sensor interface {
	fmt.Stringer

	Sensor() global.Sensor
	Read() (reading.Reading, error)
}

var sensors []Sensor

func RegisterSensor(newSensor Sensor) {
	sensors = append(sensors, newSensor)
	global.RegisterSensor(newSensor.Sensor())
}

func Sensors() []Sensor {
	return sensors
}
