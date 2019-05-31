package frontend

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	StaticFilesUrlPrefix string // likely "/static/"
	StaticFilesPath      string // likely "client/build/static/"
	FrontendBasePath     string // likely "client/build/"
}

func (s *Server) Init(router *mux.Router) {
	staticFileServer := http.FileServer(http.Dir(s.StaticFilesPath))
	router.PathPrefix(s.StaticFilesUrlPrefix).Handler(http.StripPrefix(s.StaticFilesUrlPrefix, staticFileServer))

	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, s.FrontendBasePath)
	}
	router.PathPrefix("/").HandlerFunc(indexHandler)
}
