package api

import (
	"net/http"

	"github.com/parrotmac/littleblue/pkg/internal/db"
	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/httputils"
)

type BuildConfigRouter struct {
	StorageService *db.Storage
}

func (router *BuildConfigRouter) CreateBuildConfigHandler(w http.ResponseWriter, r *http.Request) {
	buildCfg := &entities.BuildConfiguration{}
	err := httputils.ReadJsonBodyToEntity(r.Body, buildCfg)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = router.StorageService.CreateBuildConfiguration(buildCfg)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "created")
}
