package api

import (
	"net/http"

	"github.com/parrotmac/littleblue/pkg/internal/httputils"
	"github.com/parrotmac/littleblue/pkg/internal/services"
	"github.com/parrotmac/littleblue/pkg/internal/storage"
)

type BuildConfigRouter struct {
	BuildConfigService services.BuildConfigurationService
}

func (router *BuildConfigRouter) CreateBuildConfigHandler(w http.ResponseWriter, r *http.Request) {
	buildCfg := &storage.BuildConfiguration{}
	err := httputils.ReadJsonBodyToEntity(r.Body, buildCfg)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = router.BuildConfigService.CreateBuildConfiguration(buildCfg)
	if err != nil {
		httputils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httputils.RespondWithStatus(w, http.StatusCreated, "created")
}
