package ws

import (
	"TemperatureTracker/cli"
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
}

func (router Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	conn, err := config.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		var err error

		messageType, wsReader, err := conn.NextReader()
		if err != nil {
			return
		}

		wsWriter, err := conn.NextWriter(messageType)
		if err != nil {
			return
		}

		cliInterface := cli.New(router.Storage, wsReader, wsWriter)
		cliInterface.Handle()

		err = wsWriter.Close()
		if err != nil {
			return
		}
	}
}
