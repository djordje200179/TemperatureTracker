package handlers

import (
	"TemperatureTracker/sensor"
	"TemperatureTracker/storage"
	"TemperatureTracker/templates"
	"net/http"
)

type Context struct {
	Storage storage.Storage
	Sensors []sensor.Sensor

	Templates templates.Map
}

func MakeContext(storage storage.Storage) (Context, error) {
	templates, err := templates.Load()
	if err != nil {
		return Context{}, err
	}

	sensors, err := sensor.Sensors()
	if err != nil {
		return Context{}, err
	}

	return Context{Storage: storage, Sensors: sensors, Templates: templates}, nil
}

func (context Context) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", context.Index)
}
