package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/httputils"
)

func (s *apiServer) CreateBuildConfigHandler(w http.ResponseWriter, r *http.Request) {
	// FIXME: Validate user auth
	maybeRepoID := mux.Vars(r)["repo_id"]
	repoID, err := strconv.Atoi(maybeRepoID)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	buildCfg := &entities.BuildConfiguration{}
	err = httputils.ReadJsonBodyToEntity(r.Body, buildCfg)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	buildCfg.SourceRepositoryID = uint(repoID)

	err = s.Storage.CreateBuildConfiguration(buildCfg)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "created")
}

func (s *apiServer) ListRepoBuildConfigurationsHandler(w http.ResponseWriter, r *http.Request) {
	// FIXME: Validate user auth
	maybeRepoID := mux.Vars(r)["repo_id"]
	repoID, err := strconv.Atoi(maybeRepoID)
	if err != nil {
		httputils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	buildConfigs, err := s.Storage.ListRepoBuildConfigurations(uint(repoID))
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithJSON(w, http.StatusOK, buildConfigs)
}
