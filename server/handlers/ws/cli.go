package ws

import (
	"TemperatureTracker/cli"
	"net/http"
)

func (router *Router) CLI(writer http.ResponseWriter, request *http.Request) {
	conn := convertRequest(writer, request)

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
