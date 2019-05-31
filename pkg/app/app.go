package app

import (
	"github.com/gorilla/mux"
	"github.com/parrotmac/littleblue/pkg/internal/api"
	"github.com/parrotmac/littleblue/pkg/internal/builder"
	"github.com/parrotmac/littleblue/pkg/internal/db"
	"github.com/parrotmac/littleblue/pkg/internal/frontend"
	"github.com/parrotmac/littleblue/pkg/internal/httputils"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

type App struct {
	Config  *Config
	storage *db.Storage
	builder *builder.Builder

	httpServer http.Server
}

func (a *App) setupHttpServer() {
	router := mux.NewRouter()

	// TODO: Only expose some type of controller to the server, not individual pieces
	api.NewServer("/api", router, a.storage, a.builder)

	// TODO: Make frontend optional; paths configurable
	fe := &frontend.Server{
		StaticFilesUrlPrefix: "/static/",
		StaticFilesPath:      "client/build/static/",
		FrontendBasePath:     "client/build/",
	}
	fe.Init(router)
	a.httpServer = httputils.SetupServer(a.Config.ServerPort, router)
}

func (a *App) Run() {
	dataStore, err := db.Setup(a.Config.PostgresConfig)
	if err != nil {
		logrus.Fatalln(err)
	}

	a.storage = dataStore
	a.storage.AutoMigrateModels()

	// FIXME REMOVEME
	q := &builder.ChannelQueue{}
	q.Init()
	b := &builder.Builder{
		TaskQueue: q,
		Storage:   a.storage,
	}
	go b.Run()
	a.builder = b
	// FIXME /REMOVEME

	a.setupHttpServer()

	log.Printf("Starting HTTP server at %v", a.httpServer.Addr)
	log.Fatal(a.httpServer.ListenAndServe())
}
