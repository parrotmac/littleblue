package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	gitconfig "gopkg.in/src-d/go-git.v4/config"

	"github.com/parrotmac/littleblue/pkg/internal"
	"github.com/parrotmac/littleblue/pkg/internal/api"
	"github.com/parrotmac/littleblue/pkg/internal/config"
	"github.com/parrotmac/littleblue/pkg/internal/db"
	"github.com/parrotmac/littleblue/pkg/internal/docker"
	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/httputils"
	"github.com/parrotmac/littleblue/pkg/internal/meta"
)

type RefSpec gitconfig.RefSpec

// Deprecated
type EnvSettings struct {
	githubWebhookSecret string
	githubAuthToken     string

	bitbucketWebhookSecret string
	bitbucketAuthToken     string

	dockerRegistryURL      string
	dockerRegistryUsername string
	dockerRegistryPassword string
}

type MessageLevel string

const (
	MsgLevelDebug MessageLevel = "DEBUG"
	MsgLevelInfo  MessageLevel = "INFO"
	MsgLevelWarn  MessageLevel = "WARN"
	MsgLevelError MessageLevel = "ERROR"
)

// Deprecated
type Message struct {
	Level           MessageLevel `json:"level"`
	RepoFullName    string       `json:"repo_full_name"`
	BuildIdentifier string       `json:"build_identifier"`
	Body            string       `json:"body"`
}

// Deprecated
type GitRepository struct {
	FullName   string  `json:"full_name"`   // parrotmac/littleblue
	DashedName string  `json:"dashed_name"` // parrotmac-littleblue
	RepoName   string  `json:"repo_name"`   // littleblue
	GitRefSpec RefSpec `json:"ref_spec"`    // refs/heads/master

	// TODO: Refactor path scheme to be robust, supporting different providers and branches
	FilesystemPath string // workdir/repos/parrotmac-littleblue
}

type DockerBuildSpec struct {
	RegistryURL      string
	RegistryUsername string
	RegistryPassword string
	ImageName        string
	Tag              string
}

type BuildContext struct {
	Source           GitRepository
	Docker           DockerBuildSpec
	Messages         []Message
	broadcastChannel *chan Message
	BuildIdentifier  string
}

func (bCtx *BuildContext) addMessage(level MessageLevel, iface interface{}, shouldMarshal bool) {
	newMessage := Message{
		Level:           level,
		RepoFullName:    bCtx.Source.FullName,
		BuildIdentifier: bCtx.BuildIdentifier,
	}

	if shouldMarshal {
		messageJsonBytes, err := json.Marshal(iface)
		if err != nil {
			log.Println(err)
			return
		}
		newMessage.Body = string(messageJsonBytes)
	} else {
		newMessage.Body = fmt.Sprintf("%v", iface)
	}

	bCtx.Messages = append(bCtx.Messages, newMessage)

	go func() {
		*bCtx.broadcastChannel <- newMessage
	}()
}

type App struct {
	config     *config.AppConfig
	BaseRouter *mux.Router

	storage *db.Storage

	buildContexts []*BuildContext

	buildMessages meta.BuildMessageChannel

	wsClients   map[*websocket.Conn]bool
	wsBroadcast chan Message
	jobQueue    meta.JobQueue
}

func (a *App) Run() {
	server := httputils.SetupServer(a.config.ServerPort, a.BaseRouter)
	log.Printf("Starting HTTP server at %v", server.Addr)
	log.Fatal(server.ListenAndServe())
}

// FIXME: Refactor to remove all the self-referential nonsense
func NewDefaultApp() *App {
	ctx := context.Background()

	// Load configuration
	config := &config.AppConfig{}
	err := config.LoadConfig(".")
	if err != nil {
		logrus.Fatalln(err)
	}
	err = config.Validate()
	if err != nil {
		logrus.Fatalln(err)
	}

	a := App{
		BaseRouter:    mux.NewRouter(),
		config:        config,
		buildContexts: []*BuildContext{},
	}

	dataStore, err := db.Setup(config.PostgresConfig)
	if err != nil {
		logrus.Fatalln(err)
	}
	a.storage = dataStore
	a.storage.AutoMigrateModels()

	a.wsClients = make(map[*websocket.Conn]bool)
	a.wsBroadcast = make(chan Message)
	a.jobQueue = make(meta.JobQueue, 100)
	a.buildMessages = make(meta.BuildMessageChannel, 1000)

	go a.handleMessages()

	apiServer := api.Server{
		APIRouter: a.BaseRouter.PathPrefix("/api").Subrouter(),
		Storage:   a.storage,
		JobQueue:  a.jobQueue,
	}
	apiServer.Init()

	dockerClient, err := client.NewEnvClient()
	if err != nil {
		logrus.Fatalln("failed to setup docker client", logrus.WithError(err))
	}

	builder := docker.NewBuilder(&docker.BuilderConfig{
		Client: dockerClient,
	})

	// FIXME: Move this somewhere more appropriate, or remove entirely once WS streaming is reimplemented
	go func(msgs meta.BuildMessageChannel) {
		for msg := range msgs {
			if msg.DockerMessage != nil {
				logrus.Infoln(msg.BuildJobID, msg.Stage, *msg.DockerMessage)
				msgBody, err := json.Marshal(msg.DockerMessage)
				if err != nil {
					logrus.Warnln("failed to re-marshal message")
					continue
				}

				err = a.storage.AppendLog(msg.BuildJobID, entities.BuildLogKindBuild, string(msgBody))
				if err != nil {
					logrus.Error("failed to append build log", err)
				}

				continue
			}
			logrus.Infoln(msg.BuildJobID, msg.Stage, msg.PlainMessage)
		}
	}(a.buildMessages)

	workerMgr := internal.NewWorkerManager(&internal.WorkerConfig{
		StorageService: a.storage,
		JobQueue:       &a.jobQueue,
		Builder:        builder,
		BuildMessages:  a.buildMessages,
	})
	go workerMgr.Run(ctx)

	a.BaseRouter.HandleFunc("/ws", a.websocketConnectionHandler)
	initializeFrontendRoutes(a.BaseRouter)

	return &a
}
