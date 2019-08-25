package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	StaticFilesUrlPrefix string // likely "/static/"
	StaticFilesPath      string // likely "client/build/static/"
	FrontendBasePath     string // likely "client/build/"
}

func (s *Server) Init(router *gin.RouterGroup) {
	router.Static(s.StaticFilesUrlPrefix, s.StaticFilesPath)

	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, s.FrontendBasePath)
	}
	router.GET("/", gin.WrapF(indexHandler))
}
