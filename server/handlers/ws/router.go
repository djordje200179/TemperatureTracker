package ws

import (
	"TemperatureTracker/data/storage"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var config = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Router struct {
	Storage storage.Storage

	http.ServeMux
}

func NewRouter(storage storage.Storage) *Router {
	router := &Router{
		Storage: storage,
	}

	router.attachRoutes()

	return router
}

func (router *Router) attachRoutes() {
	router.HandleFunc("/cli", router.CLI)
	router.HandleFunc("/sensor", router.Sensor)
}

func convertRequest(writer http.ResponseWriter, request *http.Request) *websocket.Conn {
	connection, err := config.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err)
	}

	return connection
}
