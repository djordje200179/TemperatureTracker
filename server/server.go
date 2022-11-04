package server

import (
	"TemperatureTracker/data/storage"
	"TemperatureTracker/server/handlers/api"
	"TemperatureTracker/server/handlers/pages"
	"fmt"
	"net/http"
)

func Start(storage storage.Storage, port uint16) error {
	global := http.NewServeMux()

	pagesContext := pages.Context{Storage: storage}
	global.Handle("/", pagesContext.Handler())

	apiContext := api.Context{Storage: storage}
	global.Handle("/api/", apiContext.Handler())

	staticFilesHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("server/static")))
	global.Handle("/static/", staticFilesHandler)

	addr := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(addr, global)
	if err != nil {
		return err
	}

	fmt.Println("Listening at", addr)

	return nil
}
