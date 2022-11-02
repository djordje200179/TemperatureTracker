package sensor

import (
	"fmt"
	"time"
)

type Reading struct {
	Sensor
	Temperature
	time.Time
}

func (reading Reading) String() string {
	formattedTime := reading.Time.Format("02.01.2006. 15:04:05")
	return fmt.Sprintf("%s (%s, %s)", reading.Temperature, reading.Sensor, formattedTime)
}
