package remote

import (
	"TemperatureTracker/data/reading"
	"github.com/gorilla/websocket"
	"strconv"
	"time"
)

type Sensor struct {
	Name string

	Conn *websocket.Conn
}

func (sensor Sensor) String() string {
	return sensor.Name
}

func (sensor Sensor) Read() (reading.Reading, error) {
	err := sensor.Conn.WriteMessage(websocket.TextMessage, []byte("READ"))
	if err != nil {
		return reading.Reading{}, err
	}

	_, rawTemperature, err := sensor.Conn.ReadMessage()
	if err != nil {
		return reading.Reading{}, err
	}

	_, rawHumidity, err := sensor.Conn.ReadMessage()
	if err != nil {
		return reading.Reading{}, err
	}

	temp, err := strconv.ParseFloat(string(rawTemperature), 64)
	if err != nil {
		return reading.Reading{}, err
	}

	humidity, err := strconv.ParseFloat(string(rawHumidity), 64)
	if err != nil {
		return reading.Reading{}, err
	}

	return reading.Reading{
		Sensor:      sensor,
		Time:        time.Now(),
		Temperature: reading.Temperature(temp),
		Humidity:    reading.Humidity(humidity),
	}, nil
}
