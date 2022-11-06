package pages

import (
	"TemperatureTracker/data/storage"
	"TemperatureTracker/data/storage/cache"
	"TemperatureTracker/server/templates"
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
