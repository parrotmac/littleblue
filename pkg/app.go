package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"gopkg.in/src-d/go-git.v4/config"
)

type RefSpec config.RefSpec

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

type Message struct {
	Level           MessageLevel `json:"level"`
	RepoFullName    string       `json:"repo_full_name"`
	BuildIdentifier string       `json:"build_identifier"`
	Body            string       `json:"body"`
}

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
	config    *appConfig
	Router    *mux.Router
	APIRouter *mux.Router

	buildContexts []*BuildContext

	wsClients   map[*websocket.Conn]bool
	wsBroadcast chan Message
}

func (a *App) InitializeRouting() {
	a.Router = mux.NewRouter()
	a.Router.StrictSlash(true)
	a.APIRouter = a.Router.PathPrefix("/api").Subrouter()

	log.Print("[INIT] Setting up routes")
	a.initializeApiRoutes()
	a.initializeFrontendRoutes()

	log.Print("[INIT] Initialization complete")
}

func (a *App) initializeApiRoutes() {
	a.Router.HandleFunc("/ws", a.websocketConnectionHandler)

	a.APIRouter.HandleFunc("/jobs", a.getJobsRoute).Methods("GET")
	webhookRouter := a.APIRouter.PathPrefix("/webhook").Subrouter()
	webhookRouter.HandleFunc("", a.webhookUpdate).Methods("POST")
}

func (a *App) initializeFrontendRoutes() {
	fs := http.FileServer(http.Dir("static/"))
	a.Router.PathPrefix("/").Handler(fs)
}

func (a *App) Run() {
	bindAddress := fmt.Sprintf("0.0.0.0:%d", a.config.ServerPort)
	log.Printf("Starting HTTP server at %v", bindAddress)
	log.Fatal(http.ListenAndServe(bindAddress, a.Router))
}

func NewDefaultApp(config *appConfig) *App {
	a := App{
		config:        config,
		buildContexts: []*BuildContext{},
	}

	a.wsClients = make(map[*websocket.Conn]bool)
	a.wsBroadcast = make(chan Message)

	go a.handleMessages()

	a.InitializeRouting()

	return &a
}
