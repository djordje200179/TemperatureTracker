package ds18b20

import (
	"log"
	"os"
	"strings"
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

func findId() string {
	deviceFiles, err := os.ReadDir("/sys/bus/w1/devices")
	if err != nil {
		log.Fatal(err)
	}

	for _, deviceFile := range deviceFiles {
		if !strings.HasPrefix(deviceFile.Name(), "28-") {
			continue
		}

		return deviceFile.Name()
	}

	log.Fatal("No DS18B20 sensor found")
	return ""
}
