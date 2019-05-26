package api

import (
	"net/http"

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
