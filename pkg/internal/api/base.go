package api

import (
	"github.com/parrotmac/littleblue/pkg/internal/db"
	"log"

	"github.com/gorilla/mux"
)

type Server struct {
	APIRouter *mux.Router
	Storage   *db.Storage
}

func (s *Server) Init() {

	log.Print("[INIT] Setting up routes")

	s.initUserRoutes()
	s.initSourceProviderRoutes()
	s.initSourceRepoRoutes()
	s.initBuildConfigRoutes()

	log.Print("[INIT] Initialization complete")
}

func (s *Server) initUserRoutes() {
	userRouter := UserRouter{
		StorageService: s.Storage,
	}

	s.APIRouter.HandleFunc("/users/", userRouter.CreateUserHandler).Methods("POST")
	s.APIRouter.HandleFunc("/users/{user_id}/", userRouter.GetUserHandler).Methods("GET")
	s.APIRouter.HandleFunc("/users/{user_id}/", userRouter.UpdateUserHandler).Methods("PATCH")
}

func (s *Server) initSourceProviderRoutes() {
	router := SourceProviderRouter{
		StorageService: s.Storage,
	}
	s.APIRouter.HandleFunc("/source-providers/", router.CreateSourceProviderHandler).Methods("POST")
}

func (s *Server) initSourceRepoRoutes() {
	router := SourceRepositoryRouter{
		StorageService: s.Storage,
	}
	s.APIRouter.HandleFunc("/repos/", router.CreateSourceRepositoryHandler).Methods("POST")
}

func (s *Server) initBuildConfigRoutes() {
	router := BuildConfigRouter{
		StorageService: s.Storage,
	}
	s.APIRouter.HandleFunc("/build-configs/", router.CreateBuildConfigHandler).Methods("POST")
}
