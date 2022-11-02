package global

type Sensor string

func (sensor Sensor) String() string {
	return string(sensor)
}
