package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/parrotmac/littleblue/pkg/internal/db"
	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/httputils"
)

type BuildJobRouter struct {
	StorageService *db.Storage
}

func (router *BuildJobRouter) CreateBuildJobHandler(w http.ResponseWriter, r *http.Request) {
	buildJob := &entities.BuildJob{}
	err := httputils.ReadJsonBodyToEntity(r.Body, buildJob)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = router.StorageService.CreateBuildJob(buildJob)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "created")
}

func (router *BuildJobRouter) WebhookJobHandler(w http.ResponseWriter, r *http.Request) {
	repoUuid := mux.Vars(r)["repo_uuid"]

	_, err := router.StorageService.LookupWebhookRepoConfigurations(repoUuid)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO: Add build(s) to work queue

	httputils.RespondWithStatus(w, http.StatusOK, "ok")
}
