package pages

import (
	"TemperatureTracker/data/storage"
	"TemperatureTracker/data/storage/cache"
	"TemperatureTracker/server/templates"
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

	router.HandleFunc("/", router.Index)

	return router
}

func returnPage(templateName string, writer http.ResponseWriter, data any) {
	writer.Header().Set("Content-Type", "text/html")

	tmpl := templates.Get(templateName)

	err := tmpl.Execute(writer, data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
