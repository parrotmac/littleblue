package api

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/parrotmac/littleblue/pkg/internal/db"
	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/httputils"
	"github.com/parrotmac/littleblue/pkg/internal/uuidgen"
)

type SourceRepositoryRouter struct {
	StorageService *db.Storage
}

func (router *SourceRepositoryRouter) CreateSourceRepositoryHandler(w http.ResponseWriter, r *http.Request) {
	sourceRepo := &entities.SourceRepository{}

	err := httputils.ReadJsonBodyToEntity(r.Body, sourceRepo)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	// Overwrite on create
	sourceRepo.RepoUUID = uuidgen.NewUndashed()
	err = router.StorageService.CreateSourceRepository(sourceRepo)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "created")
}

func (router *SourceRepositoryRouter) ListSourceRepositoriesHandler(w http.ResponseWriter, r *http.Request) {
	hardcodedUserID := uint(1)

	sourceRepos, err := router.StorageService.ListUserSourceRepositories(hardcodedUserID)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	for i := range sourceRepos {
		sourceRepos[i].AuthenticationCodeSecret = ""
	}

	httputils.RespondWithJSON(w, http.StatusOK, sourceRepos)
}

func (router *SourceRepositoryRouter) ListRepoBuildJobsHandler(w http.ResponseWriter, r *http.Request) {
	hardcodedUserID := uint(1)
	repoID, ok := mux.Vars(r)["repo_id"]
	repoIntID, err := strconv.Atoi(repoID)
	if !ok || err != nil {
		httputils.RespondWithError(w, http.StatusBadRequest, errors.New("bad URL"))
		return
	}

	buildJobs, err := router.StorageService.ListBuildJobsForRepoAndUserID(hardcodedUserID, uint(repoIntID))
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	httputils.RespondWithJSON(w, http.StatusOK, buildJobs)
}
