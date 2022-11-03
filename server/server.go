package server

import (
	"TemperatureTracker/data/storage"
	"TemperatureTracker/server/handlers"
	"fmt"
	"net/http"
)

func Start(storage storage.Storage, port uint16) error {
	indexMux := http.NewServeMux()

	indexContext := handlers.Context{Storage: storage}
	indexContext.RegisterHandlers(indexMux)

	addr := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(addr, indexMux)
	if err != nil {
		return err
	}

	return nil
}
