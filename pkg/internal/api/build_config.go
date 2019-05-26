package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/parrotmac/littleblue/pkg/internal/db"
	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/httputils"
)

type BuildConfigRouter struct {
	StorageService *db.Storage
}

func (router *BuildConfigRouter) CreateBuildConfigHandler(w http.ResponseWriter, r *http.Request) {
	// FIXME: Validate user auth
	maybeRepoID := mux.Vars(r)["repo_id"]
	repoID, err := strconv.Atoi(maybeRepoID)

	buildCfg := &entities.BuildConfiguration{}
	err = httputils.ReadJsonBodyToEntity(r.Body, buildCfg)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	buildCfg.SourceRepositoryID = uint(repoID)

	err = router.StorageService.CreateBuildConfiguration(buildCfg)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "created")
}

func (router *BuildConfigRouter) ListRepoBuildConfigurationsHandler(w http.ResponseWriter, r *http.Request) {
	// FIXME: Validate user auth
	maybeRepoID := mux.Vars(r)["repo_id"]
	repoID, err := strconv.Atoi(maybeRepoID)
	if err != nil {
		httputils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	buildConfigs, err := router.StorageService.ListRepoBuildConfigurations(uint(repoID))
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithJSON(w, http.StatusOK, buildConfigs)
}
