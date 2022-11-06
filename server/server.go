package server

import (
	"TemperatureTracker/data/storage"
	"TemperatureTracker/server/handlers/api"
	"TemperatureTracker/server/handlers/pages"
	"TemperatureTracker/server/handlers/ws"
	"fmt"
	"net/http"
)

type routers struct {
	pages *pages.Router
	ws    ws.Router
	api   *api.Router

	files http.Handler
}

type Server struct {
	http.ServeMux
	routers
}

func New(storage storage.Storage) *Server {
	server := &Server{
		routers: routers{
			pages: pages.NewRouter(storage),
			ws:    ws.Router{Storage: storage},
			api:   api.NewRouter(storage),
			files: http.FileServer(http.Dir("server/static")),
		},
	}

	server.attachRoutes()

	return server
}

func (server *Server) attachRoutes() {
	server.Handle("/", server.routers.pages)
	server.Handle("/ws", http.StripPrefix("/ws", server.routers.ws))
	server.Handle("/api/", http.StripPrefix("/api", server.routers.api))
	server.Handle("/static/", http.StripPrefix("/static/", server.routers.files))
}

func (server *Server) Start(port uint16) error {
	addr := fmt.Sprintf(":%d", port)

	err := http.ListenAndServe(addr, server)
	if err != nil {
		return err
	}

	return nil
}
