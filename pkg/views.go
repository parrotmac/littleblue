package pkg

import (
	"net/http"

	"github.com/gorilla/mux"
)

func initializeFrontendRoutes(router *mux.Router) {

	staticUrlPrefix := "/static/"
	clientDirectoryPath := "client/build/static/"

	staticFileServer := http.FileServer(http.Dir(clientDirectoryPath))
	router.PathPrefix(staticUrlPrefix).Handler(http.StripPrefix(staticUrlPrefix, staticFileServer))

	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "client/build/")
	}
	router.PathPrefix("/").HandlerFunc(indexHandler)
}
