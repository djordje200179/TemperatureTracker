package api

import (
	"TemperatureTracker/data/storage"
	"TemperatureTracker/data/storage/cache"
	"encoding/json"
	"net/http"
)

type Router struct {
	storage storage.Storage
	cache   storage.Cache

	http.ServeMux
}

func NewRouter(storage storage.Storage) http.Handler {
	router := &Router{
		storage: storage,
		cache:   cache.Instance(),
	}

	router.HandleFunc("/latest", router.Latest)

	return router
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
