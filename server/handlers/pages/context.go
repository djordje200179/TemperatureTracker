package pages

import (
	"TemperatureTracker/data/storage"
	"net/http"
)

type Context struct {
	Storage storage.Storage
}

func (context Context) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", context.Index)

	return mux
}
