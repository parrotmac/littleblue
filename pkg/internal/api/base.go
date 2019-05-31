package api

import (
	"github.com/gorilla/mux"
	"github.com/parrotmac/littleblue/pkg/internal/builder"

	"github.com/parrotmac/littleblue/pkg/internal/db"
)

type server struct {
	APIRouter *mux.Router
	Storage   *db.Storage
	Builder   *builder.Builder
}

func NewServer(pathPrefix string, router *mux.Router, storage *db.Storage, builder *builder.Builder) *server {
	s := &server{
		APIRouter: router.PathPrefix(pathPrefix).Subrouter(),
		Storage:   storage,
		Builder:   builder,
	}

	s.initUserRoutes()
	s.initSourceProviderRoutes()
	s.initSourceRepoRoutes()
	s.initBuildJobRoutes()
	return s
}

func (s *server) initUserRoutes() {
	userRouter := UserRouter{
		StorageService: s.Storage,
	}

	s.APIRouter.HandleFunc("/users/", userRouter.CreateUserHandler).Methods("POST")
	s.APIRouter.HandleFunc("/users/{user_id}/", userRouter.GetUserHandler).Methods("GET")
	s.APIRouter.HandleFunc("/users/{user_id}/", userRouter.UpdateUserHandler).Methods("PATCH")
}

func (s *server) initSourceProviderRoutes() {
	router := SourceProviderRouter{
		StorageService: s.Storage,
	}
	s.APIRouter.HandleFunc("/source-providers/", router.CreateSourceProviderHandler).Methods("POST")
	s.APIRouter.HandleFunc("/source-providers/", router.ListSourceProvidersHandler).Methods("GET")
}

func (s *server) initSourceRepoRoutes() {
	repoRouter := SourceRepositoryRouter{
		StorageService: s.Storage,
	}
	repoSubrouter := s.APIRouter.PathPrefix("/repos/").Subrouter()

	repoSubrouter.HandleFunc("/", repoRouter.CreateSourceRepositoryHandler).Methods("POST")
	repoSubrouter.HandleFunc("/", repoRouter.ListSourceRepositoriesHandler).Methods("GET")

	configRouter := BuildConfigRouter{
		StorageService: s.Storage,
	}
	repoSubrouter.HandleFunc("/{repo_id}/configs/", configRouter.CreateBuildConfigHandler).Methods("POST")
	repoSubrouter.HandleFunc("/{repo_id}/configs/", configRouter.ListRepoBuildConfigurationsHandler).Methods("GET")
}

func (s *server) initBuildJobRoutes() {
	router := BuildJobRouter{
		StorageService: s.Storage,
		Builder:        s.Builder,
	}
	s.APIRouter.HandleFunc("/jobs/", router.CreateBuildJobHandler).Methods("POST")

	s.APIRouter.HandleFunc("/webhook/{repo_uuid}/", router.WebhookJobHandler).Methods("POST")
}
