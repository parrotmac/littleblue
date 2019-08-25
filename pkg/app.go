package littleblue

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/parrotmac/littleblue/pkg/internal/api"
	"github.com/parrotmac/littleblue/pkg/internal/builder"
	"github.com/parrotmac/littleblue/pkg/internal/db"
	"github.com/parrotmac/littleblue/pkg/internal/frontend"
	"github.com/parrotmac/littleblue/pkg/internal/httputils"
)

type App struct {
	Config  *Config
	storage *db.Storage
	builder *builder.Builder

	httpServer http.Server
}

func (a *App) setupHttpServer() {
	router := gin.Default()

	// TODO: Only expose some type of controller to the server, not individual pieces
	apiGroup := router.Group("/api")
	api.NewServer(apiGroup, a.storage, a.builder)

	// TODO: Make frontend optional; paths configurable
	fe := &frontend.Server{
		StaticFilesUrlPrefix: "/static/",
		StaticFilesPath:      "client/build/static/",
		FrontendBasePath:     "client/build/",
	}
	fe.Init(router.Group("/"))
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
