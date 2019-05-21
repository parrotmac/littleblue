package api

import (
	"net/http"

	"github.com/parrotmac/littleblue/pkg/internal/httputils"
	"github.com/parrotmac/littleblue/pkg/internal/services"
)

type BuildConfigRouter struct {
	BuildConfigService services.BuildConfigurationService
}

func (router *BuildConfigRouter) CreateBuildConfigHandler(w http.ResponseWriter, r *http.Request) {
	httputils.RespondWithError(w, http.StatusNotImplemented, "not yet implemented")
}
