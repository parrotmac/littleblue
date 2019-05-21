package api

import (
	"net/http"

	"github.com/parrotmac/littleblue/pkg/internal/httputils"
	"github.com/parrotmac/littleblue/pkg/internal/services"
)

type SourceProviderRouter struct {
	SourceProviderService services.SourceProviderService
}

func (router *SourceProviderRouter) CreateSourceProviderHandler(w http.ResponseWriter, r *http.Request) {
	httputils.RespondWithError(w, http.StatusNotImplemented, "not yet implemented")
}
