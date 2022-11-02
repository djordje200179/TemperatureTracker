package local

import (
	"TemperatureTracker/data/reading"
	"TemperatureTracker/data/sensors/global"
	"os"
	"strconv"
	"strings"
	"time"
)

func Init() {
	//var err error
	//
	//err = exec.Command("bash", "-c", "modprobe w1-gpio").Run()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//err = exec.Command("bash", "-c", "modprobe w1-therm").Run()
	//if err != nil {
	//	log.Fatal(err)
	//}
}

type DS18B20 struct {
	Id string
}

func RegisterDS18B20() error {
	deviceFiles, err := os.ReadDir("/sys/bus/w1/devices")
	if err != nil {
		return err
	}

	for _, deviceFile := range deviceFiles {
		if !strings.HasPrefix(deviceFile.Name(), "28-") {
			continue
		}

		sensor := DS18B20{Id: deviceFile.Name()}
		RegisterSensor(sensor)

		break
	}

	return nil
}

func (sensor DS18B20) String() string {
	return "local/DS18B20"
}

func (sensor DS18B20) Sensor() global.Sensor {
	return global.Sensor(sensor.String())
}

func (sensor DS18B20) Read() (reading.Reading, error) {
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
