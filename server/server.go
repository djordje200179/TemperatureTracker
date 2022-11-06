package server

import (
	"TemperatureTracker/data/storage"
	"TemperatureTracker/server/handlers/api"
	"TemperatureTracker/server/handlers/pages"
	"TemperatureTracker/server/handlers/ws"
	"fmt"
	"net/http"
)

func Start(storage storage.Storage, port uint16) error {
	globalRouter := http.NewServeMux()

	pagesRouter := pages.NewRouter(storage)
	wsRouter := ws.Router{Storage: storage}
	apiRouter := api.NewRouter(storage)
	staticFilesRouter := http.FileServer(http.Dir("server/static"))

	globalRouter.Handle("/", pagesRouter)
	globalRouter.Handle("/ws", http.StripPrefix("/ws", wsRouter))
	globalRouter.Handle("/api/", http.StripPrefix("/api", apiRouter))
	globalRouter.Handle("/static/", http.StripPrefix("/static/", staticFilesRouter))

	addr := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(addr, globalRouter)
	if err != nil {
		return err
	}

	fmt.Println("Listening at", addr)

	return nil
}
