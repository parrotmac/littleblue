package main

import (
	"os"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type EnvSettings struct {
	githubWebhookSecret string
	githubAuthToken		string

	gitBranch           string

	dockerRegistryURL		string
	dockerRegistryUsername	string
	dockerRegistryPassword	string

	slackWebhookURL			string
}

type WebhookRepo struct {

}

type App struct {
	Router  *mux.Router
	APIRouter  *mux.Router
	AppSettings *EnvSettings
	Repos		map[string]WebhookRepo
	GHWebhook	*GithubWebhookRequest
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
	webhookRouter := a.APIRouter.PathPrefix("/webhook").Subrouter()
	webhookRouter.HandleFunc("", a.webhookUpdate).Methods("POST")
}

func (a *App) initializeFrontendRoutes() {
	a.Router.HandleFunc("/", a.frontendRoute).Methods("GET")
}

func (a *App) Run(addr string) {
	log.Printf("Starting HTTP server at %v", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func main() {
	a := App{}

	a.AppSettings = &EnvSettings{
		githubWebhookSecret: 	os.Getenv("GH_WEBHOOK_SECRET"),
		githubAuthToken: 		os.Getenv("GH_AUTH_TOKEN"),

		gitBranch:	os.Getenv("GIT_BRANCH"),

		dockerRegistryURL: os.Getenv("DOCKER_REGISTRY_URL"),
		dockerRegistryUsername: os.Getenv("DOCKER_REGISTRY_USER"),
		dockerRegistryPassword: os.Getenv("DOCKER_REGISTRY_PASS"),

		slackWebhookURL:	os.Getenv("SLACK_WEBHOOK_URL"),
	}

	a.InitializeRouting()

	a.Run("0.0.0.0:9000")
}
