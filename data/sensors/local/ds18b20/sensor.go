package ds18b20

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/sensors/global"
	"os"
	"strconv"
	"strings"
	"time"
)

type ds18b20 struct {
	Id string
}

func (sensor ds18b20) String() string {
	return "local/ds18b20"
}

func (sensor ds18b20) Sensor() global.Sensor {
	return global.Sensor(sensor.String())
}

func (sensor ds18b20) Read() (reading.Reading, error) {
	deviceFilePath := "/sys/bus/w1/devices/" + sensor.Id + "/w1_slave"
	deviceFileContent, err := os.ReadFile(deviceFilePath)
	if err != nil {
		return reading.Reading{}, err
	}
	deviceFileRawData := strings.Split(string(deviceFileContent), "\n")

	rawTemp, err := strconv.ParseFloat(deviceFileRawData[1][len(deviceFileRawData[1])-5:], 64)
	if err != nil {
		return reading.Reading{}, err
	}
	rawTemp /= 1000

	newReading := reading.Reading{
		Sensor: sensor.Sensor(),
		Time:   time.Now(),

		Temperature: reading.Temperature(rawTemp),
		Humidity:    -1,
	}

	return newReading, nil
}
