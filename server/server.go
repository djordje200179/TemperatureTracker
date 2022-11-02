package server

import (
	"TemperatureTracker/data/storage"
	"TemperatureTracker/server/handlers"
	"fmt"
	"net/http"
)

func Start(storage storage.Storage, port uint16) error {
	mux := http.NewServeMux()

	context := handlers.MakeContext(storage)
	context.RegisterHandlers(mux)

	addr := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		return err
	}

	return nil
}
