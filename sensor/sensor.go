package sensor

import (
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

type Sensor string

func Sensors() ([]Sensor, error) {
	deviceFiles, err := os.ReadDir("/sys/bus/w1/devices")
	if err != nil {
		return nil, err
	}

	sensors := make([]Sensor, len(deviceFiles)-1)
	for i, device := range deviceFiles {
		if i == len(deviceFiles)-1 {
			break
		}

		sensors[i] = Sensor(device.Name())
	}

	return sensors, nil
}

func (sensor Sensor) String() string {
	return string(sensor)
}

func (sensor Sensor) Read() (Reading, error) {
	deviceFilePath := "/sys/bus/w1/devices/" + sensor.String() + "/w1_slave"
	deviceFileContent, err := os.ReadFile(deviceFilePath)
	if err != nil {
		return Reading{}, err
	}
	rawDeviceFileData := strings.Split(string(deviceFileContent), "\n")

	rawTemp, err := strconv.ParseFloat(rawDeviceFileData[1][len(rawDeviceFileData[1])-5:], 64)
	if err != nil {
		return Reading{}, err
	}
	rawTemp /= 1000

	return Reading{sensor, Temperature(rawTemp), time.Now()}, nil
}
