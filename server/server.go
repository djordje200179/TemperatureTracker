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

	indexMux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	addr := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(addr, indexMux)
	if err != nil {
		return err
	}

	fmt.Println("Listening at", addr)

	return nil
}
