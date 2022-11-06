package api

import (
	"TemperatureTracker/data/storage"
	"TemperatureTracker/data/storage/cache"
	"encoding/json"
	"net/http"
)

type Router struct {
	Storage storage.Storage
	Cache   storage.Cache

	http.ServeMux
}

func NewRouter(storage storage.Storage) *Router {
	router := &Router{
		Storage: storage,
		Cache:   cache.Instance(),
	}

	router.attachRoutes()

	return router
}

func (router *Router) attachRoutes() {
	router.HandleFunc("/latest", router.Latest)
}

func returnJson(writer http.ResponseWriter, data any) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(jsonData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
