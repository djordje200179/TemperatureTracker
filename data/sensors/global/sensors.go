package global

var sensors []Sensor

func RegisterSensor(newSensor Sensor) {
	sensors = append(sensors, newSensor)
}

func Sensors() []Sensor {
	return sensors
}
