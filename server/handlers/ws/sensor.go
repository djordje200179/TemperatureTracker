package ws

import (
	"TemperatureTracker/data/logger"
	"TemperatureTracker/data/sensors/remote"
	"net/http"
	"time"
)

func (router *Router) Sensor(writer http.ResponseWriter, request *http.Request) {
	conn := convertRequest(writer, request)

	_, rawName, err := conn.ReadMessage()
	if err != nil {
		return
	}
	name := string(rawName)

	sensor := remote.Sensor{
		Name: name,
		Conn: conn,
	}

	logger.Start(sensor, router.Storage, 1*time.Minute)
}
