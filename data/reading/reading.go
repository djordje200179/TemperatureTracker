package reading

import (
	"TemperatureTracker/data/sensors/global"
	"fmt"
	"time"
)

type Reading struct {
	global.Sensor

	time.Time

	Temperature
	Humidity
}

func (reading Reading) String() string {
	//formattedTime := reading.Time.Format("02.01.2006. 15:04:05")
	return fmt.Sprintf("%s, %s", reading.Temperature, reading.Humidity)
}
