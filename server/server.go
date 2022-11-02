package server

import (
	"TemperatureTracker/server/handlers"
	"TemperatureTracker/storage"
	"fmt"
	"net/http"
)

func Start(storage storage.Storage, port uint16) error {
	var err error

	context, err := handlers.MakeContext(storage)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	context.RegisterHandlers(mux)

	addr := fmt.Sprintf(":%d", port)
	err = http.ListenAndServe(addr, mux)
	if err != nil {
		return err
	}

	return nil
}
